package storage

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/models"
	"github.com/stretchr/testify/assert"
)

func TestWeatherStorage_PostWeatherData(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	db := sqlx.NewDb(sqlDB, "sqlmock")

	w := NewWeatherStorage(db)

	type testArgs struct {
		input models.WeatherResponse
		agent models.Agent
	}

	type mockBehavior func(args testArgs, agentId int)

	testTale := []struct {
		name           string
		mockBehavior   mockBehavior
		args           testArgs
		expectedId     int
		expectingError bool
	}{
		{
			name: "Correct",
			mockBehavior: func(args testArgs, agentId int) {
				mock.ExpectBegin()

				//?
				//selectRows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				//selectRows := sqlmock.NewRows([]string{"agent_name"}).AddRow(1)

				mock.ExpectExec(regexp.QuoteMeta("SELECT agent_name FROM agents WHERE id=1")).WithArgs(args.agent.Id)
				mock.ExpectQuery("INSERT INTO requests").WithArgs(args.agent.AgentName).WillReturnRows(rows)
				mock.ExpectExec("INSERT INTO states").WithArgs(args.input.Location, args.input.Temperature,
					args.input.Pressure, args.input.Rain, args.input.Cloud, args.input.WindSpeed).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("INSERT INTO links").WithArgs(1, 1, agentId).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

			},
			args: testArgs{
				input: models.WeatherResponse{
					Temperature: 100,
					Pressure:    20,
					Rain:        12,
					Cloud:       12,
					WindSpeed:   12,
					Humidity:    12,
					Location:    "Epifan",
				},
				agent: models.Agent{
					Id:   1,
					Name: "test",
				},
			},
			expectedId: 1,
		},
	}

	for _, test := range testTale {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehavior(test.args, test.args.agent.Id)

			actualId, err := w.PostWeatherData(test.args.agent.Id, test.args.input)
			if test.expectingError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedId, actualId)
			}
		})
	}
}
