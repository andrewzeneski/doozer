package proto

type ReqGet struct {
    Path string
}

type ReqSet struct {
    Path, Body, Cas string
}

type ReqDel struct {
    Path, Cas string
}

// e.g. join 4eec5bfb.38c24ce9 1.2.3.4:999
type ReqJoin struct {
    Who, Addr string
}

type ReqCheckin struct {
    Sid, Cas string
}

type ResJoin struct {
    Seqn, Snapshot string
}

type ResCheckin struct {
    T, Cas string
}