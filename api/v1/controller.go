package v1

import (
	"errors"
	"fmt"

	"log"

	//serv "go-recommendation-system/services"
	"go-recommendation-system/protos"
	service "go-recommendation-system/services/v1"

	"github.com/beego/beego/v2/server/web"

	"go-recommendation-system/utils"
	"net/http"
)

type RegCtrl struct {
	web.Controller
}

func (ctrl *RegCtrl) Get() {
	email := ctrl.Ctx.Input.Query("email")
	pass := ctrl.Ctx.Input.Query("pass")

	if email == "" || pass == "" {
		err := errors.New("RegCtrl: failed to get email '" + email + "' or pass '" + pass + "'\n")
		ctrl.CustomAbort(http.StatusBadRequest, err.Error())
		log.Printf("%v", err)
		return
	}

	if !utils.CheckValidEmail(email) {
		err := errors.New("RegCtrl: invalid email of '" + email + "'\n")
		ctrl.CustomAbort(http.StatusBadRequest, err.Error())
		log.Printf("%v", err)
		return
	}

	if e := utils.CheckValidPassword(pass); e != nil {
		err := fmt.Errorf("RegCtrl: invalid password: err = %v\n", e)
		ctrl.CustomAbort(http.StatusBadRequest, err.Error())
		log.Printf("%v", err)
		return
	}

	rep, err1 := service.Reg(&protos.RegReq{Email: email, Password: pass})

	if err1 != nil {
		ctrl.CustomAbort(http.StatusInternalServerError, err1.Error())
		log.Printf("%v", err1)
	} else {
		ctrl.Ctx.ResponseWriter.WriteHeader(int(rep.Status))
		ctrl.Data["json"] = rep
		ctrl.ServeJSON()
		log.Printf("Rep Reg %v\n", rep)
	}
}

type ConfirmCtrl struct {
	web.Controller
}

func (ctrl *ConfirmCtrl) Get() {
	token := ctrl.Ctx.Input.Query("token")

	if token == "" {
		err := errors.New("ConfirmCtrl: failed to get token '" + token + "'\n")
		ctrl.CustomAbort(http.StatusBadRequest, err.Error())
		log.Printf("%v", err)
		return
	}

	rep, err1 := service.Confirm(&protos.ConfirmReq{Email: token})

	if err1 != nil {
		ctrl.CustomAbort(http.StatusInternalServerError, err1.Error())
		log.Printf("%v", err1)
	} else {
		ctrl.Ctx.ResponseWriter.WriteHeader(int(rep.Status))
		ctrl.Data["json"] = rep
		ctrl.ServeJSON()
		log.Printf("Rep Confirm %v\n", rep)
	}
}

type AuthCtrl struct {
	web.Controller
}

func (ctrl *AuthCtrl) Get() {
	email := ctrl.Ctx.Input.Query("email")
	pass := ctrl.Ctx.Input.Query("pass")

	if email == "" || pass == "" {
		err := errors.New("AuthCtrl: failed to get email '" + email + "' or pass '" + pass + "'\n")
		ctrl.CustomAbort(http.StatusBadRequest, err.Error())
		log.Printf("%v", err)
		return
	}

	if !utils.CheckValidEmail(email) {
		err := errors.New("RegCtrl: invalid email of '" + email + "'\n")
		ctrl.CustomAbort(http.StatusBadRequest, err.Error())
		log.Printf("%v", err)
		return
	}

	if e := utils.CheckValidPassword(pass); e != nil {
		err := fmt.Errorf("RegCtrl: invalid password: err = %v\n", e)
		ctrl.CustomAbort(http.StatusBadRequest, err.Error())
		log.Printf("%v", err)
		return
	}

	rep, err1 := service.Auth(&protos.AuthReq{Email: email, Password: pass})

	if err1 != nil {
		ctrl.Ctx.ResponseWriter.WriteHeader(int(rep.Status))
		ctrl.CustomAbort(http.StatusInternalServerError, err1.Error())
		log.Printf("%v", err1)
	} else {
		ctrl.Data["json"] = rep
		ctrl.ServeJSON()
		log.Printf("Rep Auth %v\n", rep)
	}
}

type GetRecommendationCtrl struct {
	web.Controller
}

func (ctrl *GetRecommendationCtrl) Get() {
	token := ctrl.Ctx.Input.Query("token")

	if token == "" {
		err := errors.New("GetRecommendationCtrl: failed to get token '" + token + "'\n")
		ctrl.CustomAbort(http.StatusBadRequest, err.Error())
		log.Printf("%v", err)
		return
	}

	rep, err1 := service.GetRecommendation(&protos.GetRecommendationReq{Token: token})

	if err1 != nil {
		ctrl.CustomAbort(http.StatusInternalServerError, err1.Error())
		log.Printf("%v", err1)
	} else {
		ctrl.Ctx.ResponseWriter.WriteHeader(int(rep.Status))
		ctrl.Data["json"] = rep
		ctrl.ServeJSON()
		log.Printf("Rep GetRecommendation items = %v\n", len(rep.List))
	}
}
