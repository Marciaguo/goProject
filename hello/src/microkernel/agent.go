package microkernel

import (
	"context"
)
func NewAgent(sizeEvtBuf int) *Agent{
	agt:=Agent{
		collectors:map[string]Collector{},
		evtBuf: make(chan Event,sizeEvtBuf),
		state:
	}
}
type Collector interface{
	Init(evtReceiver EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf chan Event
	cancel context.CancelFunc
	ctx context.Context
	state int
}

func(agt *Agent) EventProcessGroutine(){
	var evtSeg [10]Event
	for{
		for i := 0; i < 10; i++ {
			select {
			case evtSeg[i] = <-agt.evtBuf:
				case <-agt.ctx.Done():
					return
			}
		}
	}
}

func(agt *Agent) OnEvent(evt Event){
	agt.evtBuf<-evt
}

func (agt *Agent) Destroy() error{
	if agt.state != Waiting{
		return WrongStateError
	}
	return agt.destoryCollectors()
}

func(agt *Agent) Stop() error{
	if agt.state!=Running{
		return WrongStateError
	}
	agt.state = Waiting
	agt.cancel()
	return agt.stopCollectors()
}