package replication

type replica struct {
	IsLeader bool

	ports []int
	port  int
}

func NewReplica(ports []int, port int) *replica {
	return &replica{ports: ports, port: port, IsLeader: false}
}

func (r *replica) GetLeaderPort() int {

}

func max(slice []int) int {
	max := -1
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}
