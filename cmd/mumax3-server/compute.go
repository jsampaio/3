package main

/*
	Compute service runs jobs on this node's GPUs, if any.
*/

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/mumax/3/httpfs"
	"github.com/mumax/3/util"
)

var (
	MumaxVersion string
	GPUs         []string
	Processes    = make(map[string]*Process)
)

type Process struct {
	*exec.Cmd
	Start     time.Time
	Out       io.WriteCloser
	OutputURL string
	GUI       string
	Killed    bool
}

func (p *Process) Host() string {
	return JobHost(p.OutputURL)
}

// Runs a compute service on this node, if GPUs are available.
// The compute service asks storage nodes for a job, runs it,
// saves results over httpfs and notifies storage when ready.
func RunComputeService() {

	if len(GPUs) == 0 {
		return
	}

	// queue of available GPU numbers
	idle := make(chan int, len(GPUs))
	for i := range GPUs {
		idle <- i
	}

	for {
		gpu := <-idle // take an available GPU
		GUIAddr := fmt.Sprint(thisHost+":", GUI_PORT+gpu)
		ID := WaitForJob() // take an available job
		go func() {

			p := NewProcess("http://"+ID, gpu, GUIAddr)

			WLock()
			Processes[ID] = p
			WUnlock()

			p.Run()

			// remove from "running" list
			WLock()
			delete(Processes, ID)
			WUnlock()

			_, err := RPCCall(JobHost(ID), "UpdateJob", ID)
			if err != nil {
				log.Println(err)
			}

			// add GPU number back to idle stack
			idle <- gpu
		}()
	}
}

func (p *Process) Duration() time.Duration { return Since(time.Now(), p.Start) }

// RPC-callable function, answers by this node's time
func WhatsTheTime(string) string {
	return time.Now().Format(time.UnixDate)
}

func AskTime(host string) time.Time {
	str, _ := RPCCall(host, "WhatsTheTime", "")
	return parseTime(str)
}

func parseTime(str string) time.Time {
	t, err := time.Parse(time.UnixDate, str)
	if err != nil {
		//log.Println("ParseTime:", str, err) // TODO
	}
	return t
}

func WaitForJob() string {
	ID := FindJob()
	for ID == "" {
		time.Sleep(2 * time.Second) // TODO: don't poll
		ID = FindJob()
	}
	//log.Println("found job", ID)
	return ID
}

func FindJob() string {

	// quickly list peers first
	RLock()
	p := make([]string, 0, len(peers))
	for addr, _ := range peers {
		p = append(p, addr)
	}
	RUnlock()

	// then do slow RPC calls without blocking the rest of the program
	for _, addr := range p {
		ID, _ := RPCCall(addr, "GiveJob", thisAddr)
		if ID != "" {
			return ID
		}
	}
	return ""
}

// RPC-callable function kills job corresponding to given job id.
// The job has to be running on this node.
func Kill(id string) string {
	log.Println("KILL", id)

	//	if JobHost(id) != thisAddr {
	//		ret, _ := RPCCall(JobHost(id), "Kill", id)
	//		return ret
	//	}

	WLock() // modifies Cmd state
	defer WUnlock()

	job := Processes[id]
	if job == nil {
		return fmt.Sprintf("kill %v: job not running.", id)
	}
	job.Killed = true
	err := job.Cmd.Process.Kill()
	if err != nil {
		return err.Error()
	}

	return "" // OK
}

// prepare exec.Cmd to run mumax3 compute process
func NewProcess(inputURL string, gpu int, webAddr string) *Process {
	// prepare command
	command := *flag_mumax
	gpuFlag := fmt.Sprint(`-gpu=`, gpu)
	httpFlag := fmt.Sprint(`-http=`, webAddr)
	cacheFlag := fmt.Sprint(`-cache=`, *flag_cachedir)
	forceFlag := `-f=0`
	cmd := exec.Command(command, gpuFlag, httpFlag, cacheFlag, forceFlag, inputURL)

	// Pipe stdout, stderr to log file over httpfs
	outDir := util.NoExt(inputURL) + ".out"
	httpfs.Mkdir(outDir)
	out, errD := httpfs.Create(outDir + "/stdout.txt")
	if errD != nil {
		log.Println("makeProcess", errD)
	}
	cmd.Stderr = out
	cmd.Stdout = out

	return &Process{Cmd: cmd, Start: time.Now(), Out: out, OutputURL: OutputDir(inputURL), GUI: webAddr}
}

func (p *Process) Run() {
	log.Println("=> exec  ", p.Path, p.Args)

	defer p.Out.Close()

	httpfs.Put(p.OutputURL+"host", []byte(thisAddr))

	startTime := AskTime(p.Host())
	httpfs.Put(p.OutputURL+"start", []byte(startTime.Format(time.UnixDate)))

	WLock()               // Cmd.Start() modifies state
	err1 := p.Cmd.Start() // err?
	WUnlock()

	timeOffset := time.Now().Sub(startTime) // our clock is most likely out-of-sync with host
	tick := time.NewTicker(KeepaliveInterval)

	// need initial alive in case watchdog sniffs between start and first alive tick
	httpfs.Put(p.OutputURL+"alive", []byte(time.Now().Add(timeOffset).Format(time.UnixDate)))
	go func() {
		for t := range tick.C {
			httpfs.Put(p.OutputURL+"alive", []byte(t.Add(timeOffset).Format(time.UnixDate)))
		}
	}()

	err2 := p.Cmd.Wait()
	tick.Stop()

	status := -1

	// TODO: determine proper status number
	if err1 != nil || err2 != nil {
		log.Println(p.Path, p.Args, err1, err2)
		status = 1
	} else {
		status = 0
	}

	if p.Killed {
		httpfs.Put(p.OutputURL+"killed", []byte(time.Now().Format(time.UnixDate)))
	} else {
		httpfs.Put(p.OutputURL+"exitstatus", []byte(fmt.Sprint(status)))
	}

	return // TODO: write stat
}

func DetectGPUs() {
	if GPUs != nil {
		panic("multiple DetectGPUs() calls")
	}

	for i := 0; i < MAXGPU; i++ {
		gpuflag := fmt.Sprint("-gpu=", i)
		out, err := exec.Command(*flag_mumax, "-test", gpuflag).Output()
		if err == nil {
			info := string(out)
			if strings.HasSuffix(info, "\n") {
				info = info[:len(info)-1]
			}
			log.Println("gpu", i, ":", info)
			GPUs = append(GPUs, info)
		}
	}
}

func DetectMumax() {
	out, err := exec.Command(*flag_mumax, "-test", "-v").CombinedOutput()
	info := string(out)
	if err == nil {
		split := strings.SplitN(info, "\n", 2)
		version := split[0]
		log.Println("have", version)
		MumaxVersion = version
	} else {
		MumaxVersion = fmt.Sprint(*flag_mumax, "-test", ": ", err, info)
	}
}
