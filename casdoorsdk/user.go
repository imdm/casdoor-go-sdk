// Copyright 2021 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casdoorsdk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ManagedAccount struct {
	Application string `xorm:"varchar(100)" json:"application"`
	Username    string `xorm:"varchar(100)" json:"username"`
	Password    string `xorm:"varchar(100)" json:"password"`
	SigninUrl   string `xorm:"varchar(200)" json:"signinUrl"`
}

// User has the same definition as https://github.com/casdoor/casdoor/blob/master/object/user.go#L24
type User struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	UpdatedTime string `xorm:"varchar(100)" json:"updatedTime"`

	Id                string   `xorm:"varchar(100) index" json:"id"`
	Type              string   `xorm:"varchar(100)" json:"type"`
	Password          string   `xorm:"varchar(100)" json:"password"`
	PasswordSalt      string   `xorm:"varchar(100)" json:"passwordSalt"`
	DisplayName       string   `xorm:"varchar(100)" json:"displayName"`
	FirstName         string   `xorm:"varchar(100)" json:"firstName"`
	LastName          string   `xorm:"varchar(100)" json:"lastName"`
	Avatar            string   `xorm:"varchar(500)" json:"avatar"`
	PermanentAvatar   string   `xorm:"varchar(500)" json:"permanentAvatar"`
	Email             string   `xorm:"varchar(100) index" json:"email"`
	EmailVerified     bool     `json:"emailVerified"`
	Phone             string   `xorm:"varchar(100) index" json:"phone"`
	Location          string   `xorm:"varchar(100)" json:"location"`
	Address           []string `json:"address"`
	Affiliation       string   `xorm:"varchar(100)" json:"affiliation"`
	Title             string   `xorm:"varchar(100)" json:"title"`
	IdCardType        string   `xorm:"varchar(100)" json:"idCardType"`
	IdCard            string   `xorm:"varchar(100) index" json:"idCard"`
	Homepage          string   `xorm:"varchar(100)" json:"homepage"`
	Bio               string   `xorm:"varchar(100)" json:"bio"`
	Tag               string   `xorm:"varchar(100)" json:"tag"`
	Region            string   `xorm:"varchar(100)" json:"region"`
	Language          string   `xorm:"varchar(100)" json:"language"`
	Gender            string   `xorm:"varchar(100)" json:"gender"`
	Birthday          string   `xorm:"varchar(100)" json:"birthday"`
	Education         string   `xorm:"varchar(100)" json:"education"`
	Score             int      `json:"score"`
	Karma             int      `json:"karma"`
	Ranking           int      `json:"ranking"`
	IsDefaultAvatar   bool     `json:"isDefaultAvatar"`
	IsOnline          bool     `json:"isOnline"`
	IsAdmin           bool     `json:"isAdmin"`
	IsGlobalAdmin     bool     `json:"isGlobalAdmin"`
	IsForbidden       bool     `json:"isForbidden"`
	IsDeleted         bool     `json:"isDeleted"`
	SignupApplication string   `xorm:"varchar(100)" json:"signupApplication"`
	Hash              string   `xorm:"varchar(100)" json:"hash"`
	PreHash           string   `xorm:"varchar(100)" json:"preHash"`

	CreatedIp      string `xorm:"varchar(100)" json:"createdIp"`
	LastSigninTime string `xorm:"varchar(100)" json:"lastSigninTime"`
	LastSigninIp   string `xorm:"varchar(100)" json:"lastSigninIp"`

	GitHub   string `xorm:"github varchar(100)" json:"github"`
	Google   string `xorm:"varchar(100)" json:"google"`
	QQ       string `xorm:"qq varchar(100)" json:"qq"`
	WeChat   string `xorm:"wechat varchar(100)" json:"wechat"`
	Facebook string `xorm:"facebook varchar(100)" json:"facebook"`
	DingTalk string `xorm:"dingtalk varchar(100)" json:"dingtalk"`
	Weibo    string `xorm:"weibo varchar(100)" json:"weibo"`
	Gitee    string `xorm:"gitee varchar(100)" json:"gitee"`
	LinkedIn string `xorm:"linkedin varchar(100)" json:"linkedin"`
	Wecom    string `xorm:"wecom varchar(100)" json:"wecom"`
	Lark     string `xorm:"lark varchar(100)" json:"lark"`
	Gitlab   string `xorm:"gitlab varchar(100)" json:"gitlab"`
	Adfs     string `xorm:"adfs varchar(100)" json:"adfs"`
	Baidu    string `xorm:"baidu varchar(100)" json:"baidu"`
	Alipay   string `xorm:"alipay varchar(100)" json:"alipay"`
	Casdoor  string `xorm:"casdoor varchar(100)" json:"casdoor"`
	Infoflow string `xorm:"infoflow varchar(100)" json:"infoflow"`
	Apple    string `xorm:"apple varchar(100)" json:"apple"`
	AzureAD  string `xorm:"azuread varchar(100)" json:"azuread"`
	Slack    string `xorm:"slack varchar(100)" json:"slack"`
	Steam    string `xorm:"steam varchar(100)" json:"steam"`
	Bilibili string `xorm:"bilibili varchar(100)" json:"bilibili"`
	Okta     string `xorm:"okta varchar(100)" json:"okta"`
	Douyin   string `xorm:"douyin varchar(100)" json:"douyin"`
	Custom   string `xorm:"custom varchar(100)" json:"custom"`

	//WebauthnCredentials []webauthn.Credential `xorm:"webauthnCredentials blob" json:"webauthnCredentials"`

	Ldap       string            `xorm:"ldap varchar(100)" json:"ldap"`
	Properties map[string]string `json:"properties"`

	Roles       []*Role       `json:"roles"`
	Permissions []*Permission `json:"permissions"`

	LastSigninWrongTime string `xorm:"varchar(100)" json:"lastSigninWrongTime"`
	SigninWrongTimes    int    `json:"signinWrongTimes"`

	ManagedAccounts []ManagedAccount `xorm:"managedAccounts blob" json:"managedAccounts"`
}

func (c *Client) GetUsers() ([]*User, error) {
	queryMap := map[string]string{
		"owner": c.cfg.OrganizationName,
	}

	url := c.GetUrl("get-users", queryMap)

	bytes, err := c.DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var users []*User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) GetSortedUsers(sorter string, limit int) ([]*User, error) {
	queryMap := map[string]string{
		"owner":  c.cfg.OrganizationName,
		"sorter": sorter,
		"limit":  strconv.Itoa(limit),
	}

	url := c.GetUrl("get-sorted-users", queryMap)

	bytes, err := c.DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var users []*User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) GetPaginationUsers(p int, pageSize int, queryMap map[string]string) ([]*User, int, error) {
	queryMap["owner"] = c.cfg.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-users", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	if response.Status != "ok" {
		return nil, 0, fmt.Errorf(response.Msg)
	}

	bytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var users []*User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, 0, err
	}
	return users, int(response.Data2.(float64)), nil
}

func (c *Client) GetUserCount(isOnline string) (int, error) {
	queryMap := map[string]string{
		"owner":    c.cfg.OrganizationName,
		"isOnline": isOnline,
	}

	url := c.GetUrl("get-user-count", queryMap)

	bytes, err := c.DoGetBytesRaw(url)
	if err != nil {
		return -1, err
	}

	var count int
	err = json.Unmarshal(bytes, &count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (c *Client) GetUser(name string) (*User, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.cfg.OrganizationName, name),
	}

	url := c.GetUrl("get-user", queryMap)

	bytes, err := c.DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) GetUserByEmail(email string) (*User, error) {
	queryMap := map[string]string{
		"owner": c.cfg.OrganizationName,
		"email": email,
	}

	url := c.GetUrl("get-user", queryMap)

	bytes, err := c.DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) GetUserByPhone(phone string) (*User, error) {
	queryMap := map[string]string{
		"owner": c.cfg.OrganizationName,
		"phone": phone,
	}

	url := c.GetUrl("get-user", queryMap)

	bytes, err := c.DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) GetUserByUserId(userId string) (*User, error) {
	queryMap := map[string]string{
		"owner":  c.cfg.OrganizationName,
		"userId": userId,
	}

	url := c.GetUrl("get-user", queryMap)

	bytes, err := c.DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// note: oldPassword is not required, if you don't need, just pass a empty string
func (c *Client) SetPassword(owner, name, oldPassword, newPassword string) (bool, error) {
	param := map[string]string{
		"userOwner":   owner,
		"userName":    name,
		"oldPassword": oldPassword,
		"newPassword": newPassword,
	}

	bytes, err := json.Marshal(param)
	if err != nil {
		return false, err
	}

	resp, err := c.DoPost("set-password", nil, bytes, true, false)
	if err != nil {
		return false, err
	}

	return resp.Status == "ok", nil
}

func (c *Client) UpdateUserById(id string, user *User) (bool, error) {
	_, affected, err := c.modifyUserById("update-user", id, user, nil)
	return affected, err
}

func (c *Client) UpdateUser(user *User) (bool, error) {
	_, affected, err := c.modifyUser("update-user", user, nil)
	return affected, err
}

func (c *Client) UpdateUserForColumns(user *User, columns []string) (bool, error) {
	_, affected, err := c.modifyUser("update-user", user, columns)
	return affected, err
}

func (c *Client) AddUser(user *User) (bool, error) {
	_, affected, err := c.modifyUser("add-user", user, nil)
	return affected, err
}

func (c *Client) DeleteUser(user *User) (bool, error) {
	_, affected, err := c.modifyUser("delete-user", user, nil)
	return affected, err
}

func (c *Client) CheckUserPassword(user *User) (bool, error) {
	response, _, err := c.modifyUser("check-user-password", user, nil)
	return response.Status == "ok", err
}

func (u User) GetId() string {
	return fmt.Sprintf("%s/%s", u.Owner, u.Name)
}
