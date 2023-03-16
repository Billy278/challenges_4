package member

type Member interface {
	GetMember(id int) (MemberImpl, error)
}
