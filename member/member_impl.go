package member

import "errors"

type MemberImpl struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Reasson   string
}

var Data = map[int]MemberImpl{
	1: {
		Nama:      "Billy",
		Alamat:    "jln jamin ginting, berastagi",
		Pekerjaan: "Pelajar",
		Reasson:   "Karena Peformanya yang unggul serta mudah digunakan",
	},
	2: {
		Nama:      "Bima",
		Alamat:    "jln jamin ginting, pajak roga",
		Pekerjaan: "Pelajar",
		Reasson:   "Karena lagi booming dan banyak di butuhkan",
	},
	3: {
		Nama:      "Rima",
		Alamat:    "jln jamin ginting, kabanjahe",
		Pekerjaan: "Pelajar",
		Reasson:   "karena populer",
	},
}

func (m *MemberImpl) GetMember(id int) (MemberImpl, error) {
	if res, ok := Data[id]; ok {
		return res, nil

	}
	return MemberImpl{}, errors.New("tidak ditemukan")

}
