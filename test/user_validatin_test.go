package test

import (
	"example.com/sa-67-example/entity"
	"fmt"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestStudentID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`student_id is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: "",
			FirstName: "สมชาย",
			LastName:  "สาย",
			Email:     "a@gmail.com",
			Phone:     "0612223333",
			Profile:   "",
			Link:      "https://www.linkedin.com/company/ilink/",
			Password:  "",
			BirthDay:  time.Now(),
			GenderID:1,
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})

	t.Run(`StudentID not pattern is not true`, func(t *testing.T) {
		user2 := entity.User{
			StudentID: "G6508463",
			FirstName: "สมชาย",
			LastName:  "สาย",
			Email:     "a@gmail.com",
			Phone:     "0612223333",
			Profile:   "",
			Link:      "https://www.linkedin.com/company/ilink/",
			Password:  "",
			BirthDay:  time.Now(),
			GenderID:1,
		}
		ok, err := govalidator.ValidateStruct(user2)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", user2.StudentID)))
	})
}

func TestPhone (t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Phone is required",func(t *testing.T){
		user := entity.User{
			StudentID: "B6508463",
			FirstName: "สมชาย",
			LastName:  "สาย",
			Email:     "a@gmail.com",
			Phone:     "",
			Profile:   "",
			Link:      "https://www.linkedin.com/company/ilink/",
			Password:  "",
			BirthDay:  time.Now(),
			GenderID:1,
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is required"))

	})

	t.Run("Phone pattern is not true",func(t *testing.T){
		user := entity.User{
			StudentID: "B6508463",
			FirstName: "สมชาย",
			LastName:  "สาย",
			Email:     "a@gmail.com",
			Phone:     "061222333",
			Profile:   "",
			Link:      "https://www.linkedin.com/company/ilink/",
			Password:  "",
			BirthDay:  time.Now(),
			GenderID:1,
		}
		ok,err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as stringlength(10|10)", user.Phone)))
	})
}

func TestEmail (t *testing.T){
	g := NewGomegaWithT(t)
	t.Run("Email is required",func(t *testing.T){
		user := entity.User{
			StudentID: "B6508463",
			FirstName: "สมชาย",
			LastName:  "สาย",
			Email:     "",
			Phone:     "0612223333",
			Profile:   "",
			Link:      "https://www.linkedin.com/company/ilink/",
			Password:  "",
			BirthDay:  time.Now(),
			GenderID:1,
		}
		ok,err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run("Email Pattern Is Not True", func(t *testing.T){
		user := entity.User{
			StudentID: "B6508463",
			FirstName: "สมชาย",
			LastName:  "สาย",
			Email:     "a@@gmail.com",
			Phone:     "0612223333",
			Profile:   "",
			Link:      "https://www.linkedin.com/company/ilink/",
			Password:  "",
			BirthDay:  time.Now(),
			GenderID:1,
		}
		ok,err := govalidator.ValidateStruct(user)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		fmt.Println(err.Error())
		g.Expect(err.Error()).To(Equal("Email is invalid"))
	})
}

func TestAllTrue (t *testing.T){
	g:=NewGomegaWithT(t)
	t.Run("All True",func(t *testing.T){
		user := entity.User{
			StudentID: "B6508463",
			FirstName: "สมชาย",
			LastName:  "สาย",
			Email:     "a@gmail.com",
			Phone:     "0612223333",
			Profile:   "",
			Link:      "https://www.linkedin.com/company/ilink/",
			Password:  "",
			BirthDay:  time.Now(),
			GenderID:1,
		}
		ok,err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}