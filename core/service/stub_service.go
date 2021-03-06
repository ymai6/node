/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package service

import (
	"github.com/mysteriumnetwork/node/identity"
	dto_discovery "github.com/mysteriumnetwork/node/service_discovery/dto"
	"github.com/mysteriumnetwork/node/session"
)

var _ Service = &serviceFake{}

type serviceFake struct {
	onStartReturnError error
}

func (service *serviceFake) Start(identity.Identity) (dto_discovery.ServiceProposal, session.ConfigProvider, error) {
	return dto_discovery.ServiceProposal{}, nil, service.onStartReturnError
}

func (service *serviceFake) Wait() error {
	return nil
}

func (service *serviceFake) Stop() error {
	return nil
}
