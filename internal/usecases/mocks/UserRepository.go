// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	domain "GraphQLTestCase/internal/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, entity
func (_m *UserRepository) Create(ctx context.Context, entity domain.User) error {
	ret := _m.Called(ctx, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, limit, offset
func (_m *UserRepository) GetAll(ctx context.Context, limit int, offset int) ([]*domain.User, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]*domain.User, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*domain.User); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*domain.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIds provides a mock function with given fields: ctx, ids
func (_m *UserRepository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]*domain.User, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) ([]*domain.User, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) []*domain.User); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uuid.UUID) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, entity
func (_m *UserRepository) Update(ctx context.Context, entity domain.User) error {
	ret := _m.Called(ctx, entity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) error); ok {
		r0 = rf(ctx, entity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
