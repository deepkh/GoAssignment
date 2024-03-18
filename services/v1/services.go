package v1

import (
	"context"
	"fmt"
	"go-recommendation-system/db"
	"go-recommendation-system/protos"
	"go-recommendation-system/utils"
	"log"
	"net/http"
	"time"
)

type Services struct {
	protos.UnimplementedRegServServer
	protos.UnimplementedAuthServServer
	protos.UnimplementedRecommendationServServer
}

var defaultTokenExpiredHour = time.Hour * 24 * 14

func handErr(err *error, s string, rs *string) {
	if *err == nil {
		return
	}

	// assign the error status and the reason
	l := fmt.Sprintf("Req %v failed, err = %v", s, *err)
	log.Printf("%s", l)
	*rs = l

	// It's not necessary to pass the error to the upstream Grpc Server,
	// due to the status code and the reason is already assigned
	*err = nil
}

func NewServices() *Services {
	return &Services{}
}

func (p *Services) Reg(ctx context.Context, in *protos.RegReq) (r *protos.RegRep, err error) {
	r = &protos.RegRep{}
	defer handErr(&err, "Reg", &r.Reason)

	// Check user if exist
	var u *protos.User
	u, err = db.QueryUser(in.Email)
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	// Bad Request
	if u != nil {
		err = fmt.Errorf("User %v already created.", in.Email)
		r.Status = http.StatusBadRequest
		return
	}

	// create a user
	err = db.CreateUser(&protos.User{Email: in.Email, PasswordHashed: utils.PasswordHash(in.Password), Confirm: 0, Timestamp: time.Now().UnixMicro()})
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	// generate a user token
	var token string
	token, err = utils.GenereUserToken(in.Email, defaultTokenExpiredHour)
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	r.Status = http.StatusOK
	r.Reason = ""
	r.ConfirmLink = fmt.Sprintf("http://127.0.0.1:8888/v1/confirm?token=%v", token)
	log.Printf("Req Reg %v ", in)
	return
}

func (p *Services) Confirm(ctx context.Context, in *protos.ConfirmReq) (r *protos.ConfirmRep, err error) {
	r = &protos.ConfirmRep{}
	defer handErr(&err, "Confirm", &r.Reason)

	// verify & decrypt
	var email string = ""
	email, err = utils.ParseUserToken(in.Email)
	if err != nil {
		r.Status = http.StatusBadRequest
		return
	}

	// Check user if exist
	var u *protos.User
	u, err = db.QueryUser(email)
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	// Bad Request
	if u == nil {
		err = fmt.Errorf("User not found.")
		r.Status = http.StatusBadRequest
		return
	}

	// update a user
	u.Confirm = 1
	err = db.UpdateUser(u)
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	r.Status = http.StatusOK
	r.Reason = "Confirmed!"
	log.Printf("Req Confirm in = %v, email = %v", in, email)
	return
}

func (p *Services) Auth(ctx context.Context, in *protos.AuthReq) (r *protos.AuthRep, err error) {
	r = &protos.AuthRep{}
	defer handErr(&err, "Auth", &r.Reason)

	// Check user if exist
	var u *protos.User
	u, err = db.QueryUser(in.Email)
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	// Bad Request
	if u == nil {
		err = fmt.Errorf("User %v not found.", in.Email)
		r.Status = http.StatusBadRequest
		return
	}

	// check password
	match := utils.PasswordVerify(in.Password, u.PasswordHashed)
	if !match {
		err = fmt.Errorf("User %s password is wrong.", in.Email)
		r.Status = http.StatusBadRequest
		return
	}

	// check confirm
	if u.Confirm == 0 {
		err = fmt.Errorf("User %s not confirmed.", in.Email)
		r.Status = http.StatusBadRequest
		return
	}

	// generate a user token
	var token string
	token, err = utils.GenereUserToken(u.Email, defaultTokenExpiredHour)
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	r.Status = http.StatusOK
	r.Reason = ""
	r.AuthedUser = &protos.AuthedUser{Token: token}
	log.Printf("Req Auth %v", in)
	return
}

func (p *Services) CheckToken(ctx context.Context, in *protos.CheckTokenReq) (r *protos.CheckTokenRep, err error) {
	r = &protos.CheckTokenRep{}
	defer handErr(&err, "CheckToken", &r.Reason)

	r.Status = http.StatusOK
	r.Reason = ""
	log.Printf("Req CheckToken %v", in)
	return
}

func (p *Services) GetRecommendation(ctx context.Context, in *protos.GetRecommendationReq) (r *protos.GetRecommendationRep, err error) {
	r = &protos.GetRecommendationRep{}
	defer handErr(&err, "GetRecommendation", &r.Reason)

	// verify & decrypt
	var email string = ""
	email, err = utils.ParseUserToken(in.Token)
	if err != nil {
		r.Status = http.StatusBadRequest
		return
	}

	// Check user if exist
	var u *protos.User
	u, err = db.QueryUser(email)
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	// Bad Request
	if u == nil {
		err = fmt.Errorf("User not found.")
		r.Status = http.StatusBadRequest
		return
	}

	// check confirm
	if u.Confirm == 0 {
		err = fmt.Errorf("User %s not confirmed.", u.Email)
		r.Status = http.StatusBadRequest
		return
	}

	// Send Recommendations to user
	r.List, err = db.QueryRecommendations()
	if err != nil {
		r.Status = http.StatusInternalServerError
		return
	}

	r.Status = http.StatusOK
	r.Reason = ""
	log.Printf("Req GetRecommendation %v, u = %v", in, u)
	return
}
