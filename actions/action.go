package actions

import (
	"codemi/loker/helpers"
	"codemi/loker/models"
	"codemi/loker/usecase"
	"strconv"
	"strings"
)

var storage *usecase.Storage

func Init(params []string) {
	if storage != nil {
		helpers.Println("Init sudah pernah dilakukan sebelumnya")
		return
	}
	if len(params) == 0 {
		helpers.Println("action init harus memiliki 1 parameter")
		return
	}
	size, err := strconv.Atoi(params[0])
	if err != nil {
		helpers.Println("jumlah loker harus berupa angka")
		return
	}
	storage = usecase.NewStorage(size)
	helpers.Printf("Berhasil membuat loker dengan jumlah %d \n", size)
}

func Status(params []string) {
	data := storage.GetAll()
	if len(data) == 0 {
		helpers.Println("Data loker kosong, silahkan melakukan init")
		return
	}
	helpers.Printf("No Loker\tTipe Identitas\tNo Identitas\n")
	checkType := func(data *models.Data) string {
		if data != nil {
			return data.Type
		}
		return "kosong"
	}
	checkNo := func(data *models.Data) string {
		if data != nil {
			return strconv.FormatInt(data.Number, 10)
		}
		return "kosong"
	}
	for i, row := range data {
		helpers.Printf("%d\t\t%s\t\t%s\n", i+1, checkType(row), checkNo(row))
	}
}

func Input(params []string) {
	if !defaultValidation("input", 2, params) {
		return
	}
	identityType := params[0]
	number, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		helpers.Println("nomor identitas harus berupa angka")
		return
	}
	index, err := storage.Save(identityType, number)
	if err != nil {
		helpers.Println(err.Error())
	} else {
		helpers.Printf("Kartu identitas tersimpan di loker nomor %d\n", index+1)
	}
}

func Leave(params []string) {
	if !defaultValidation("leave", 1, params) {
		return
	}
	index, err := strconv.ParseInt(params[0], 10, 64)
	if err != nil {
		helpers.Println("nomor loker harus berupa angka")
		return
	}
	err = storage.Remove(int(index) - 1)
	if err != nil {
		helpers.Println(err.Error())
	} else {
		helpers.Printf("Loker nomor %d berhasil dikosongkan\n", index)
	}
}

func Find(params []string) {
	if !defaultValidation("find", 1, params) {
		return
	}
	number, err := strconv.ParseInt(params[0], 10, 64)
	if err != nil {
		helpers.Println("nomor identitas harus berupa angka")
		return
	}
	index, err := storage.Get(number)
	if err != nil {
		helpers.Println(err.Error())
	} else {
		helpers.Printf("Kartu identitas tersebut berada di loker nomor %d\n", index+1)
	}
}

func Search(params []string) {
	if !defaultValidation("search", 1, params) {
		return
	}
	identityType := params[0]
	data := storage.GetByType(identityType)
	result := []string{}
	for _, row := range data {
		result = append(result, strconv.FormatInt(row.Number, 10))
	}
	if len(result) == 0 {
		helpers.Println("Tidak ada data dengan type ", identityType)
	} else {
		helpers.Println(strings.Join(result, ", "))
	}
}

func defaultValidation(action string, numParam int, params []string) bool {
	if storage == nil {
		helpers.Println("Silah Melakukan init terlebih dahulu")
		return false
	}
	if len(params) < numParam {
		helpers.Printf("action %s harus memiliki %d paramter\n", action, numParam)
		return false
	}
	return true
}
