package pkg

type SocketState string

const (
	SocketStateListen      SocketState = "0A" // TCP_LISTEN
	SocketStateEstablished             = "01" // TCP_ESTABLISHED
	SocketStateSynSent                 = "02" // TCP_SYN_SENT
	SocketStateSynRecv                 = "03" // TCP_SYN_RECV
	SocketStateFinWait1                = "04" // TCP_FIN_WAIT1
	SocketStateFinWait2                = "05" // TCP_FIN_WAIT2
	SocketStateTimeWait                = "06" // TCP_TIME_WAIT
)

func (s SocketState) String() string {
	switch s {
	case SocketStateListen:
		return "LISTEN"
	case SocketStateEstablished:
		return "ESTABLISHED"
	case SocketStateSynSent:
		return "SYN_SENT"
	case SocketStateSynRecv:
		return "SYN_RECV"
	case SocketStateFinWait1:
		return "FIN_WAIT1"
	case SocketStateFinWait2:
		return "FIN_WAIT2"
	case SocketStateTimeWait:
		return "TIME_WAIT"
	default:
		return string(s)
	}
}
