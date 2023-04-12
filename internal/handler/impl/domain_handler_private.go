package impl

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hmsidm/internal/domain/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (a *application) findIpaById(tx *gorm.DB, orgId string, uuid string) (data *model.Domain, err error) {
	data = &model.Domain{}
	*data, err = a.domain.repository.FindById(tx, orgId, uuid)
	if err != nil {
		return nil, err
	}
	if *data.DomainType != model.DomainTypeIpa {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Wrong domain type")
	}

	if data.IpaDomain == nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "No IPA data found for the domain")
	}
	return data, nil
}

// checkToken verifies that the provided token is valid based on the token
// stored into the database Ipa record.
// token is the token we received into the http request for PUT
// /domains/{uuid}/ipa/register
// ipa is the database ipa entity that contains the token information.
// Return nil if the check is successful, else an error with the detailed
// causes.
func (a *application) checkToken(token string, ipa *model.Ipa) error {
	if token == "" {
		return fmt.Errorf("'token' cannot be empty")
	}
	if ipa == nil {
		return fmt.Errorf("'ipa' cannot be nil")
	}
	if ipa.Token == nil || *ipa.Token == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "OTP token is required")
	}
	if ipa.TokenExpiration == nil || (*ipa.TokenExpiration == time.Time{}) {
		return echo.NewHTTPError(http.StatusBadRequest, "OTP token expiration not found")
	}
	if *ipa.Token != token {
		return echo.NewHTTPError(http.StatusBadRequest, "OTP token does not match")
	}
	if ipa.TokenExpiration.Before(time.Now()) {
		return echo.NewHTTPError(http.StatusBadRequest, "OTP token expired")
	}
	return nil
}

// existsHostInServers verify that the given fqdn exists into the list of
// IpaServer.
// fqdn is the fqdn of the host to check.
// servers is the slice of IpaServer that is defined for the IPA domain.
// Return nil if the fqdn succesfully match with some item into the list of
// servers, else return an error.
func (a *application) existsHostInServers(
	fqdn string,
	servers []model.IpaServer,
) error {
	if servers == nil {
		return fmt.Errorf("'servers' cannot be nil")
	}

	for i := range servers {
		if servers[i].FQDN == fqdn {
			return nil
		}
	}

	return fmt.Errorf("'fqdn' not found into the list of IPA servers")
}

// fillIpaDomain is a helper function to copy Ipa domain
// data between structures, to be used at register IPA domain endpoint.
// target is the destination Ipa structure, it cannot be nil.
// source is the source Ipa structure, it cannot be nil.
// Return nil if it is copied succesfully, else an error.
func (a *application) fillIpaDomain(
	target *model.Ipa,
	source *model.Ipa,
) error {
	if source == nil {
		return fmt.Errorf("'target' cannot be nil")
	}
	if target == nil {
		return fmt.Errorf("'source' cannot be nil")
	}

	target.CaCerts = source.CaCerts
	for i := range source.CaCerts {
		target.CaCerts[i].IpaID = target.ID
	}
	target.RealmDomains = source.RealmDomains
	target.Servers = source.Servers
	for i := range source.Servers {
		target.Servers[i].IpaID = target.ID
	}

	return nil
}