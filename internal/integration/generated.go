package integration

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
)

type Species string

const (
	SpeciesDog        Species = "DOG"
	SpeciesCoelacanth Species = "COELACANTH"
)

// queryWithFragmentsBeingsAnimal includes the requested fields of the GraphQL type Animal.
type queryWithFragmentsBeingsAnimal struct {
	Typename string                             `json:"__typename"`
	Id       string                             `json:"id"`
	Name     string                             `json:"name"`
	Hair     queryWithFragmentsBeingsHair       `json:"hair"`
	Species  Species                            `json:"species"`
	Owner    queryWithFragmentsBeingsOwnerBeing `json:"-"`
}

func (v *queryWithFragmentsBeingsAnimal) UnmarshalJSON(b []byte) error {

	type queryWithFragmentsBeingsAnimalWrapper queryWithFragmentsBeingsAnimal

	var firstPass struct {
		*queryWithFragmentsBeingsAnimalWrapper
		Owner json.RawMessage `json:"owner"`
	}
	firstPass.queryWithFragmentsBeingsAnimalWrapper = (*queryWithFragmentsBeingsAnimalWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.Owner
		raw := firstPass.Owner
		err = __unmarshalqueryWithFragmentsBeingsOwnerBeing(
			target, raw)
		if err != nil {
			return err
		}
	}
	return nil
}

// queryWithFragmentsBeingsBeing includes the requested fields of the GraphQL interface Being.
//
// queryWithFragmentsBeingsBeing is implemented by the following types:
// queryWithFragmentsBeingsUser
// queryWithFragmentsBeingsAnimal
//
// The GraphQL type's documentation follows.
//
//
type queryWithFragmentsBeingsBeing interface {
	implementsGraphQLInterfacequeryWithFragmentsBeingsBeing()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	GetId() string
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *queryWithFragmentsBeingsUser) implementsGraphQLInterfacequeryWithFragmentsBeingsBeing() {}

// GetTypename is a part of, and documented with, the interface queryWithFragmentsBeingsBeing.
func (v *queryWithFragmentsBeingsUser) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithFragmentsBeingsBeing.
func (v *queryWithFragmentsBeingsUser) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithFragmentsBeingsBeing.
func (v *queryWithFragmentsBeingsUser) GetName() string { return v.Name }

func (v *queryWithFragmentsBeingsAnimal) implementsGraphQLInterfacequeryWithFragmentsBeingsBeing() {}

// GetTypename is a part of, and documented with, the interface queryWithFragmentsBeingsBeing.
func (v *queryWithFragmentsBeingsAnimal) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithFragmentsBeingsBeing.
func (v *queryWithFragmentsBeingsAnimal) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithFragmentsBeingsBeing.
func (v *queryWithFragmentsBeingsAnimal) GetName() string { return v.Name }

func __unmarshalqueryWithFragmentsBeingsBeing(v *queryWithFragmentsBeingsBeing, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "User":
		*v = new(queryWithFragmentsBeingsUser)
		return json.Unmarshal(m, *v)
	case "Animal":
		*v = new(queryWithFragmentsBeingsAnimal)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for queryWithFragmentsBeingsBeing: "%v"`, tn.TypeName)
	}
}

// queryWithFragmentsBeingsHair includes the requested fields of the GraphQL type BeingsHair.
type queryWithFragmentsBeingsHair struct {
	HasHair bool `json:"hasHair"`
}

// queryWithFragmentsBeingsOwnerAnimal includes the requested fields of the GraphQL type Animal.
type queryWithFragmentsBeingsOwnerAnimal struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithFragmentsBeingsOwnerBeing includes the requested fields of the GraphQL interface Being.
//
// queryWithFragmentsBeingsOwnerBeing is implemented by the following types:
// queryWithFragmentsBeingsOwnerUser
// queryWithFragmentsBeingsOwnerAnimal
//
// The GraphQL type's documentation follows.
//
//
type queryWithFragmentsBeingsOwnerBeing interface {
	implementsGraphQLInterfacequeryWithFragmentsBeingsOwnerBeing()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	GetId() string
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *queryWithFragmentsBeingsOwnerUser) implementsGraphQLInterfacequeryWithFragmentsBeingsOwnerBeing() {
}

// GetTypename is a part of, and documented with, the interface queryWithFragmentsBeingsOwnerBeing.
func (v *queryWithFragmentsBeingsOwnerUser) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithFragmentsBeingsOwnerBeing.
func (v *queryWithFragmentsBeingsOwnerUser) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithFragmentsBeingsOwnerBeing.
func (v *queryWithFragmentsBeingsOwnerUser) GetName() string { return v.Name }

func (v *queryWithFragmentsBeingsOwnerAnimal) implementsGraphQLInterfacequeryWithFragmentsBeingsOwnerBeing() {
}

// GetTypename is a part of, and documented with, the interface queryWithFragmentsBeingsOwnerBeing.
func (v *queryWithFragmentsBeingsOwnerAnimal) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithFragmentsBeingsOwnerBeing.
func (v *queryWithFragmentsBeingsOwnerAnimal) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithFragmentsBeingsOwnerBeing.
func (v *queryWithFragmentsBeingsOwnerAnimal) GetName() string { return v.Name }

func __unmarshalqueryWithFragmentsBeingsOwnerBeing(v *queryWithFragmentsBeingsOwnerBeing, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "User":
		*v = new(queryWithFragmentsBeingsOwnerUser)
		return json.Unmarshal(m, *v)
	case "Animal":
		*v = new(queryWithFragmentsBeingsOwnerAnimal)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for queryWithFragmentsBeingsOwnerBeing: "%v"`, tn.TypeName)
	}
}

// queryWithFragmentsBeingsOwnerUser includes the requested fields of the GraphQL type User.
type queryWithFragmentsBeingsOwnerUser struct {
	Typename    string `json:"__typename"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	LuckyNumber int    `json:"luckyNumber"`
}

// queryWithFragmentsBeingsUser includes the requested fields of the GraphQL type User.
type queryWithFragmentsBeingsUser struct {
	Typename    string                       `json:"__typename"`
	Id          string                       `json:"id"`
	Name        string                       `json:"name"`
	LuckyNumber int                          `json:"luckyNumber"`
	Hair        queryWithFragmentsBeingsHair `json:"hair"`
}

// queryWithFragmentsResponse is returned by queryWithFragments on success.
type queryWithFragmentsResponse struct {
	Beings []queryWithFragmentsBeingsBeing `json:"-"`
}

func (v *queryWithFragmentsResponse) UnmarshalJSON(b []byte) error {

	type queryWithFragmentsResponseWrapper queryWithFragmentsResponse

	var firstPass struct {
		*queryWithFragmentsResponseWrapper
		Beings []json.RawMessage `json:"beings"`
	}
	firstPass.queryWithFragmentsResponseWrapper = (*queryWithFragmentsResponseWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.Beings
		raw := firstPass.Beings
		*target = make(
			[]queryWithFragmentsBeingsBeing,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			err = __unmarshalqueryWithFragmentsBeingsBeing(
				target, raw)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// queryWithInterfaceListFieldBeingsAnimal includes the requested fields of the GraphQL type Animal.
type queryWithInterfaceListFieldBeingsAnimal struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithInterfaceListFieldBeingsBeing includes the requested fields of the GraphQL interface Being.
//
// queryWithInterfaceListFieldBeingsBeing is implemented by the following types:
// queryWithInterfaceListFieldBeingsUser
// queryWithInterfaceListFieldBeingsAnimal
//
// The GraphQL type's documentation follows.
//
//
type queryWithInterfaceListFieldBeingsBeing interface {
	implementsGraphQLInterfacequeryWithInterfaceListFieldBeingsBeing()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	GetId() string
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *queryWithInterfaceListFieldBeingsUser) implementsGraphQLInterfacequeryWithInterfaceListFieldBeingsBeing() {
}

// GetTypename is a part of, and documented with, the interface queryWithInterfaceListFieldBeingsBeing.
func (v *queryWithInterfaceListFieldBeingsUser) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithInterfaceListFieldBeingsBeing.
func (v *queryWithInterfaceListFieldBeingsUser) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithInterfaceListFieldBeingsBeing.
func (v *queryWithInterfaceListFieldBeingsUser) GetName() string { return v.Name }

func (v *queryWithInterfaceListFieldBeingsAnimal) implementsGraphQLInterfacequeryWithInterfaceListFieldBeingsBeing() {
}

// GetTypename is a part of, and documented with, the interface queryWithInterfaceListFieldBeingsBeing.
func (v *queryWithInterfaceListFieldBeingsAnimal) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithInterfaceListFieldBeingsBeing.
func (v *queryWithInterfaceListFieldBeingsAnimal) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithInterfaceListFieldBeingsBeing.
func (v *queryWithInterfaceListFieldBeingsAnimal) GetName() string { return v.Name }

func __unmarshalqueryWithInterfaceListFieldBeingsBeing(v *queryWithInterfaceListFieldBeingsBeing, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "User":
		*v = new(queryWithInterfaceListFieldBeingsUser)
		return json.Unmarshal(m, *v)
	case "Animal":
		*v = new(queryWithInterfaceListFieldBeingsAnimal)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for queryWithInterfaceListFieldBeingsBeing: "%v"`, tn.TypeName)
	}
}

// queryWithInterfaceListFieldBeingsUser includes the requested fields of the GraphQL type User.
type queryWithInterfaceListFieldBeingsUser struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithInterfaceListFieldResponse is returned by queryWithInterfaceListField on success.
type queryWithInterfaceListFieldResponse struct {
	Beings []queryWithInterfaceListFieldBeingsBeing `json:"-"`
}

func (v *queryWithInterfaceListFieldResponse) UnmarshalJSON(b []byte) error {

	type queryWithInterfaceListFieldResponseWrapper queryWithInterfaceListFieldResponse

	var firstPass struct {
		*queryWithInterfaceListFieldResponseWrapper
		Beings []json.RawMessage `json:"beings"`
	}
	firstPass.queryWithInterfaceListFieldResponseWrapper = (*queryWithInterfaceListFieldResponseWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.Beings
		raw := firstPass.Beings
		*target = make(
			[]queryWithInterfaceListFieldBeingsBeing,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			err = __unmarshalqueryWithInterfaceListFieldBeingsBeing(
				target, raw)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// queryWithInterfaceListPointerFieldBeingsAnimal includes the requested fields of the GraphQL type Animal.
type queryWithInterfaceListPointerFieldBeingsAnimal struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithInterfaceListPointerFieldBeingsBeing includes the requested fields of the GraphQL interface Being.
//
// queryWithInterfaceListPointerFieldBeingsBeing is implemented by the following types:
// queryWithInterfaceListPointerFieldBeingsUser
// queryWithInterfaceListPointerFieldBeingsAnimal
//
// The GraphQL type's documentation follows.
//
//
type queryWithInterfaceListPointerFieldBeingsBeing interface {
	implementsGraphQLInterfacequeryWithInterfaceListPointerFieldBeingsBeing()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	GetId() string
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *queryWithInterfaceListPointerFieldBeingsUser) implementsGraphQLInterfacequeryWithInterfaceListPointerFieldBeingsBeing() {
}

// GetTypename is a part of, and documented with, the interface queryWithInterfaceListPointerFieldBeingsBeing.
func (v *queryWithInterfaceListPointerFieldBeingsUser) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithInterfaceListPointerFieldBeingsBeing.
func (v *queryWithInterfaceListPointerFieldBeingsUser) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithInterfaceListPointerFieldBeingsBeing.
func (v *queryWithInterfaceListPointerFieldBeingsUser) GetName() string { return v.Name }

func (v *queryWithInterfaceListPointerFieldBeingsAnimal) implementsGraphQLInterfacequeryWithInterfaceListPointerFieldBeingsBeing() {
}

// GetTypename is a part of, and documented with, the interface queryWithInterfaceListPointerFieldBeingsBeing.
func (v *queryWithInterfaceListPointerFieldBeingsAnimal) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithInterfaceListPointerFieldBeingsBeing.
func (v *queryWithInterfaceListPointerFieldBeingsAnimal) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithInterfaceListPointerFieldBeingsBeing.
func (v *queryWithInterfaceListPointerFieldBeingsAnimal) GetName() string { return v.Name }

func __unmarshalqueryWithInterfaceListPointerFieldBeingsBeing(v *queryWithInterfaceListPointerFieldBeingsBeing, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "User":
		*v = new(queryWithInterfaceListPointerFieldBeingsUser)
		return json.Unmarshal(m, *v)
	case "Animal":
		*v = new(queryWithInterfaceListPointerFieldBeingsAnimal)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for queryWithInterfaceListPointerFieldBeingsBeing: "%v"`, tn.TypeName)
	}
}

// queryWithInterfaceListPointerFieldBeingsUser includes the requested fields of the GraphQL type User.
type queryWithInterfaceListPointerFieldBeingsUser struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithInterfaceListPointerFieldResponse is returned by queryWithInterfaceListPointerField on success.
type queryWithInterfaceListPointerFieldResponse struct {
	Beings []*queryWithInterfaceListPointerFieldBeingsBeing `json:"-"`
}

func (v *queryWithInterfaceListPointerFieldResponse) UnmarshalJSON(b []byte) error {

	type queryWithInterfaceListPointerFieldResponseWrapper queryWithInterfaceListPointerFieldResponse

	var firstPass struct {
		*queryWithInterfaceListPointerFieldResponseWrapper
		Beings []json.RawMessage `json:"beings"`
	}
	firstPass.queryWithInterfaceListPointerFieldResponseWrapper = (*queryWithInterfaceListPointerFieldResponseWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.Beings
		raw := firstPass.Beings
		*target = make(
			[]*queryWithInterfaceListPointerFieldBeingsBeing,
			len(raw))
		for i, raw := range raw {
			target := &(*target)[i]
			*target = new(queryWithInterfaceListPointerFieldBeingsBeing)
			err = __unmarshalqueryWithInterfaceListPointerFieldBeingsBeing(
				*target, raw)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// queryWithInterfaceNoFragmentsBeing includes the requested fields of the GraphQL interface Being.
//
// queryWithInterfaceNoFragmentsBeing is implemented by the following types:
// queryWithInterfaceNoFragmentsBeingUser
// queryWithInterfaceNoFragmentsBeingAnimal
//
// The GraphQL type's documentation follows.
//
//
type queryWithInterfaceNoFragmentsBeing interface {
	implementsGraphQLInterfacequeryWithInterfaceNoFragmentsBeing()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	GetId() string
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *queryWithInterfaceNoFragmentsBeingUser) implementsGraphQLInterfacequeryWithInterfaceNoFragmentsBeing() {
}

// GetTypename is a part of, and documented with, the interface queryWithInterfaceNoFragmentsBeing.
func (v *queryWithInterfaceNoFragmentsBeingUser) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithInterfaceNoFragmentsBeing.
func (v *queryWithInterfaceNoFragmentsBeingUser) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithInterfaceNoFragmentsBeing.
func (v *queryWithInterfaceNoFragmentsBeingUser) GetName() string { return v.Name }

func (v *queryWithInterfaceNoFragmentsBeingAnimal) implementsGraphQLInterfacequeryWithInterfaceNoFragmentsBeing() {
}

// GetTypename is a part of, and documented with, the interface queryWithInterfaceNoFragmentsBeing.
func (v *queryWithInterfaceNoFragmentsBeingAnimal) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface queryWithInterfaceNoFragmentsBeing.
func (v *queryWithInterfaceNoFragmentsBeingAnimal) GetId() string { return v.Id }

// GetName is a part of, and documented with, the interface queryWithInterfaceNoFragmentsBeing.
func (v *queryWithInterfaceNoFragmentsBeingAnimal) GetName() string { return v.Name }

func __unmarshalqueryWithInterfaceNoFragmentsBeing(v *queryWithInterfaceNoFragmentsBeing, m json.RawMessage) error {
	if string(m) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(m, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "User":
		*v = new(queryWithInterfaceNoFragmentsBeingUser)
		return json.Unmarshal(m, *v)
	case "Animal":
		*v = new(queryWithInterfaceNoFragmentsBeingAnimal)
		return json.Unmarshal(m, *v)
	default:
		return fmt.Errorf(
			`Unexpected concrete type for queryWithInterfaceNoFragmentsBeing: "%v"`, tn.TypeName)
	}
}

// queryWithInterfaceNoFragmentsBeingAnimal includes the requested fields of the GraphQL type Animal.
type queryWithInterfaceNoFragmentsBeingAnimal struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithInterfaceNoFragmentsBeingUser includes the requested fields of the GraphQL type User.
type queryWithInterfaceNoFragmentsBeingUser struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}

// queryWithInterfaceNoFragmentsMeUser includes the requested fields of the GraphQL type User.
type queryWithInterfaceNoFragmentsMeUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// queryWithInterfaceNoFragmentsResponse is returned by queryWithInterfaceNoFragments on success.
type queryWithInterfaceNoFragmentsResponse struct {
	Being queryWithInterfaceNoFragmentsBeing  `json:"-"`
	Me    queryWithInterfaceNoFragmentsMeUser `json:"me"`
}

func (v *queryWithInterfaceNoFragmentsResponse) UnmarshalJSON(b []byte) error {

	type queryWithInterfaceNoFragmentsResponseWrapper queryWithInterfaceNoFragmentsResponse

	var firstPass struct {
		*queryWithInterfaceNoFragmentsResponseWrapper
		Being json.RawMessage `json:"being"`
	}
	firstPass.queryWithInterfaceNoFragmentsResponseWrapper = (*queryWithInterfaceNoFragmentsResponseWrapper)(v)

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		target := &v.Being
		raw := firstPass.Being
		err = __unmarshalqueryWithInterfaceNoFragmentsBeing(
			target, raw)
		if err != nil {
			return err
		}
	}
	return nil
}

// queryWithVariablesResponse is returned by queryWithVariables on success.
type queryWithVariablesResponse struct {
	User queryWithVariablesUser `json:"user"`
}

// queryWithVariablesUser includes the requested fields of the GraphQL type User.
type queryWithVariablesUser struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LuckyNumber int    `json:"luckyNumber"`
}

// simpleQueryMeUser includes the requested fields of the GraphQL type User.
type simpleQueryMeUser struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LuckyNumber int    `json:"luckyNumber"`
}

// simpleQueryResponse is returned by simpleQuery on success.
type simpleQueryResponse struct {
	Me simpleQueryMeUser `json:"me"`
}

func simpleQuery(
	ctx context.Context,
	client graphql.Client,
) (*simpleQueryResponse, error) {
	var retval simpleQueryResponse
	err := client.MakeRequest(
		ctx,
		"simpleQuery",
		`
query simpleQuery {
	me {
		id
		name
		luckyNumber
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

func queryWithVariables(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*queryWithVariablesResponse, error) {
	variables := map[string]interface{}{
		"id": id,
	}

	var retval queryWithVariablesResponse
	err := client.MakeRequest(
		ctx,
		"queryWithVariables",
		`
query queryWithVariables ($id: ID!) {
	user(id: $id) {
		id
		name
		luckyNumber
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}

func queryWithInterfaceNoFragments(
	ctx context.Context,
	client graphql.Client,
	id string,
) (*queryWithInterfaceNoFragmentsResponse, error) {
	variables := map[string]interface{}{
		"id": id,
	}

	var retval queryWithInterfaceNoFragmentsResponse
	err := client.MakeRequest(
		ctx,
		"queryWithInterfaceNoFragments",
		`
query queryWithInterfaceNoFragments ($id: ID!) {
	being(id: $id) {
		__typename
		id
		name
	}
	me {
		id
		name
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}

func queryWithInterfaceListField(
	ctx context.Context,
	client graphql.Client,
	ids []string,
) (*queryWithInterfaceListFieldResponse, error) {
	variables := map[string]interface{}{
		"ids": ids,
	}

	var retval queryWithInterfaceListFieldResponse
	err := client.MakeRequest(
		ctx,
		"queryWithInterfaceListField",
		`
query queryWithInterfaceListField ($ids: [ID!]!) {
	beings(ids: $ids) {
		__typename
		id
		name
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}

func queryWithInterfaceListPointerField(
	ctx context.Context,
	client graphql.Client,
	ids []string,
) (*queryWithInterfaceListPointerFieldResponse, error) {
	variables := map[string]interface{}{
		"ids": ids,
	}

	var retval queryWithInterfaceListPointerFieldResponse
	err := client.MakeRequest(
		ctx,
		"queryWithInterfaceListPointerField",
		`
query queryWithInterfaceListPointerField ($ids: [ID!]!) {
	beings(ids: $ids) {
		__typename
		id
		name
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}

func queryWithFragments(
	ctx context.Context,
	client graphql.Client,
	ids []string,
) (*queryWithFragmentsResponse, error) {
	variables := map[string]interface{}{
		"ids": ids,
	}

	var retval queryWithFragmentsResponse
	err := client.MakeRequest(
		ctx,
		"queryWithFragments",
		`
query queryWithFragments ($ids: [ID!]!) {
	beings(ids: $ids) {
		__typename
		id
		... on Being {
			id
			name
		}
		... on Animal {
			id
			hair {
				hasHair
			}
			species
			owner {
				__typename
				id
				... on Being {
					name
				}
				... on User {
					luckyNumber
				}
			}
		}
		... on Lucky {
			luckyNumber
		}
		... on User {
			hair {
				color
			}
		}
	}
}
`,
		&retval,
		variables,
	)
	return &retval, err
}
