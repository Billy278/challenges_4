package member

import "errors"

type MemberImpl struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Reasson   string
}

func (m *MemberImpl) GetMember(id int) (MemberImpl, error) {
	data := map[int]MemberImpl{
		1: MemberImpl{
			Nama:      "Billy",
			Alamat:    "jln jamin ginting, berastagi",
			Pekerjaan: "Pelajar",
			Reasson:   "Karena Peformanya yang unggul serta mudah digunakan",
		},
		2: MemberImpl{
			Nama:      "Bima",
			Alamat:    "jln jamin ginting, pajak roga",
			Pekerjaan: "Pelajar",
			Reasson:   "Karena lagi booming dan banyak di butuhkan",
		},
		3: MemberImpl{
			Nama:      "Rima",
			Alamat:    "jln jamin ginting, kabanjahe",
			Pekerjaan: "Pelajar",
			Reasson:   "karena populer",
		},
	}

	if res, ok := data[id]; ok {
		return res, nil

	}
	return MemberImpl{}, errors.New("tidak ditemukan")

}
