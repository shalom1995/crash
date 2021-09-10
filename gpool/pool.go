package gpool

var (
	PoolGO        *Pool
	PoolGOTokenID *Pool
	PoolGOOwner   *Pool
	PoolGOBounce  *Pool
)

//定义一个Pool协程池的类型
type Pool struct {
	EntryChannel chan *Task
	JobsChannel  chan *Task
	workerNum    int
}

func init() {
	PoolGO = NewPool(10)
	go PoolGO.Run()

	PoolGOTokenID = NewPool(40)
	go PoolGOTokenID.Run()

	PoolGOOwner = NewPool(10)
	go PoolGOOwner.Run()

	PoolGOBounce = NewPool(10)
	go PoolGOBounce.Run()
}

func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		workerNum:    cap,
	}

	return &p
}

//	执行任务队列中的任务，属于消费者
func (p *Pool) do(workerID int) {
	for task := range p.JobsChannel {
		task.Execute()
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		// 启动特定数量的协程来工作
		go p.do(i)
	}

	for task := range p.EntryChannel {
		//	将任务写入，交给do去执行
		p.JobsChannel <- task
	}
}
