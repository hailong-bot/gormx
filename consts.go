package gormx

const (
	ASC Sort = iota
	DESC
)

const (
	NoLock Lock = iota
	IS
	IX
)

type Lock uint8

type Sort uint8

func (s Sort) ToString() string {
	switch s {
	case ASC:
		return "ASC"
	case DESC:
		return "DESC"
	}
	panic("not supported to this sort")
}
