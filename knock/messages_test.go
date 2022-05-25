package knock

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

func TestMessages_List(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"items":[{"__cursor":"big-cursor","__typename":"Message","archived_at":null,"channel_id":"5da042d7-02ee-46ed-8b91-9b5717da2028","data":{"middle-name":"alfred","welcome":"to jurassic park"},"id":"message-id","inserted_at":"2022-05-17T00:34:18.277163Z","read_at":null,"recipient":"tom","seen_at":null,"source":{"__typename":"NotificationSource","key":"test","version_id":"4dae021a-ba51-473f-9038-77041da8131c"},"status":"delivered","tenant":null,"updated_at":"2022-05-17T00:34:18.318283Z","workflow":"test"}],"page_info":{"__typename":"PageInfo","after":"big-after","before":null,"page_size":1}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	have, err := client.Messages.List(ctx, &ListMessagesRequest{
		PageSize: 1,
	})

	if err != nil {
		fmt.Println(err)
	}

	want := &ListMessagesResponse{
		PageInfo: PageInfo{
			PageSize: 1,
			After:    "big-after",
		},
		Items: []*Message{
			{
				Cursor:     "big-cursor",
				ID:         "message-id",
				ChannelID:  "5da042d7-02ee-46ed-8b91-9b5717da2028",
				Recipient:  "tom",
				Workflow:   "test",
				Tenant:     "",
				Status:     "delivered",
				ReadAt:     time.Time{},
				SeenAt:     time.Time{},
				ArchivedAt: time.Time{},
				InsertedAt: time.Date(2022, time.May, 17, 00, 34, 18, 277163000, time.UTC),
				UpdatedAt:  time.Date(2022, time.May, 17, 00, 34, 18, 318283000, time.UTC),
				Source: &NotificationSource{
					Key:       "test",
					VersionID: "4dae021a-ba51-473f-9038-77041da8131c",
				},
				Data: map[string]interface{}{
					"middle-name": "alfred",
					"welcome":     "to jurassic park",
				},
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestMessages_Get(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__cursor":null,"__typename":"Message","archived_at":null,"channel_id":"5da042d7-02ee-46ed-8b91-9b5717da2028","data":{"middle-name":"alfred","welcome":"to jurassic park"},"id":"long-id","inserted_at":"2022-05-17T00:34:18.277163Z","read_at":null,"recipient":"tom","seen_at":null,"source":{"__typename":"NotificationSource","key":"test","version_id":"4dae021a-ba51-473f-9038-77041da8131c"},"status":"delivered","tenant":null,"updated_at":"2022-05-17T00:34:18.318283Z","workflow":"test"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient()

	have, err := client.Messages.Get(ctx, &GetMessageRequest{
		ID: "long-id",
	})

	if err != nil {
		fmt.Println(err)
	}

	want := &GetMessageResponse{
		Message: &Message{
			Cursor:     "",
			ID:         "long-id",
			ChannelID:  "5da042d7-02ee-46ed-8b91-9b5717da2028",
			Recipient:  "tom",
			Workflow:   "test",
			Tenant:     "",
			Status:     "delivered",
			ReadAt:     time.Time{},
			SeenAt:     time.Time{},
			ArchivedAt: time.Time{},
			InsertedAt: time.Date(2022, time.May, 17, 00, 34, 18, 277163000, time.UTC),
			UpdatedAt:  time.Date(2022, time.May, 17, 00, 34, 18, 318283000, time.UTC),
			Source: &NotificationSource{
				Key:       "test",
				VersionID: "4dae021a-ba51-473f-9038-77041da8131c",
			},
			Data: map[string]interface{}{
				"middle-name": "alfred",
				"welcome":     "to jurassic park",
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestMessages_GetEvents(t *testing.T) {

	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"items":[{"__cursor":"big-cursor","__typename":"MessageEvent","data":{},"environment_id":"1e37b486-593e-47fa-8049-16b1bab76084","id":"event-id-1","inserted_at":"2022-05-17T00:34:18.322544Z","recipient":"tom","type":"message.delivered"},{"__cursor":"big-cursor-2","__typename":"MessageEvent","data":{},"environment_id":"1e37b486-593e-47fa-8049-16b1bab76084","id":"event-id-2","inserted_at":"2022-05-17T00:34:18.315054Z","recipient":"tom","type":"message.sent"}],"page_info":{"__typename":"PageInfo","after":null,"before":null,"page_size":50}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient()

	have, err := client.Messages.GetEvents(ctx, &GetMessageEventsRequest{
		ID: "29GmBF0R3ZG5l34w06vyc3H9mDa",
	})

	if err != nil {
		fmt.Println(err)
	}

	testData := make(map[string]interface{})
	want := &GetMessageEventsResponse{
		Items: []*MessageEvent{
			{
				Cursor:        "big-cursor",
				ID:            "event-id-1",
				EnvironmentID: "1e37b486-593e-47fa-8049-16b1bab76084",
				Recipient:     "tom",
				Data:          testData,
			},
			{
				Cursor:        "big-cursor-2",
				ID:            "event-id-2",
				EnvironmentID: "1e37b486-593e-47fa-8049-16b1bab76084",
				Recipient:     "tom",
				Data:          testData,
			},
		},
		PageInfo: PageInfo{
			PageSize: 50,
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestMessages_GetActivities(t *testing.T) {

	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"items":[{"__cursor":"aaaa","__typename":"Activity","actor":null,"data":{"middle-name":"alfred","welcome":"to jurassic park"},"id":"activity-id","inserted_at":"2022-05-17T00:33:19.188022Z","recipient":{"__typename":"User","created_at":null,"id":"tom3","updated_at":"2022-05-15T20:41:55.082Z","Properties":null,"email":"nice@nice.com","foo":"bar","into":"main","name":"tom3"},"updated_at":"2022-05-17T00:33:19.188022Z"}],"page_info":{"__typename":"PageInfo","after":null,"before":null,"page_size":50}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient() //TODO remove any test client commented code

	have, err := client.Messages.GetActivities(ctx, &GetMessageActivitiesRequest{
		ID: "29Gm3nVijCsi2gsmrLLTpDDtHAE",
	})

	if err != nil {
		fmt.Println(err)
	}

	want := &GetMessageActivitiesResponse{
		Activities: []*MessageActivity{
			{
				Cursor: "aaaa",
				ID:     "activity-id",
				Data: map[string]interface{}{
					"welcome":     "to jurassic park",
					"middle-name": "alfred",
				},
				Recipient: &User{
					ID:        "tom3",
					Name:      "tom3",
					Email:     "nice@nice.com",
					UpdatedAt: ParseRFC3339Timestamp("2022-05-15T20:41:55.082Z"),
				},
			},
		},
		PageInfo: PageInfo{
			PageSize: 50,
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestMessages_GetContent(t *testing.T) { // TODO replace with real API test

	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__typename":"MessageContent","id":"264os2yt574e1T3jblzjbh7Qa69","data":{"bcc":"bcc@example.com","cc":"cc@example.com","from":{"email":"info-app@example.com","name":"Example App"},"html_body":"<p>example</p>","reply_to":"reply-to@example.com","subject":"Example Notification","text_body":"example","to":"user@example.com"},"inserted_at":"2021-04-06T12:00:00Z"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient() //TODO remove any test client commented code

	have, err := client.Messages.GetContent(ctx, &GetMessageContentRequest{
		ID: "264os2yt574e1T3jblzjbh7Qa69",
	})

	if err != nil {
		fmt.Println(err)
	}

	want := &GetMessageContentResponse{
		Content: &MessageContent{
			ID:         "264os2yt574e1T3jblzjbh7Qa69",
			InsertedAt: ParseRFC3339Timestamp("2021-04-06T12:00:00Z"),
			Data: map[string]interface{}{
				"bcc":       "bcc@example.com",
				"cc":        "cc@example.com",
				"html_body": "<p>example</p>",
				"reply_to":  "reply-to@example.com",
				"subject":   "Example Notification",
				"text_body": "example",
				"to":        "user@example.com",
				"from":      map[string]interface{}{"email": "info-app@example.com", "name": "Example App"},
			},
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestMessages_SetStatus(t *testing.T) { // TODO

	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__cursor":"big-cursor","__typename":"Message","archived_at":null,"channel_id":"5da042d7-02ee-46ed-8b91-9b5717da2028","data":{"middle-name":"alfred","welcome":"to jurassic park"},"id":"message-id","inserted_at":"2022-05-17T00:34:18.277163Z","read_at":null,"recipient":"tom","seen_at":null,"source":{"__typename":"NotificationSource","key":"test","version_id":"4dae021a-ba51-473f-9038-77041da8131c"},"status":"delivered","tenant":null,"updated_at":"2022-05-17T00:34:18.318283Z","workflow":"test"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient() //TODO remove any test client commented code

	have, err := client.Messages.SetStatus(ctx, &SetStatusRequest{
		ID:     "29GlZCBxKZ79J9Uquf4HXiUZNfH",
		Status: Archived,
	})

	if err != nil {
		fmt.Println(err)
	}

	want := &SetStatusResponse{Message: &Message{
		Cursor:     "big-cursor",
		ID:         "message-id",
		ChannelID:  "5da042d7-02ee-46ed-8b91-9b5717da2028",
		Recipient:  "tom",
		Workflow:   "test",
		Tenant:     "",
		Status:     "delivered",
		ReadAt:     time.Time{},
		SeenAt:     time.Time{},
		ArchivedAt: time.Time{},
		InsertedAt: time.Date(2022, time.May, 17, 00, 34, 18, 277163000, time.UTC),
		UpdatedAt:  time.Date(2022, time.May, 17, 00, 34, 18, 318283000, time.UTC),
		Source: &NotificationSource{
			Key:       "test",
			VersionID: "4dae021a-ba51-473f-9038-77041da8131c",
		},
		Data: map[string]interface{}{
			"middle-name": "alfred",
			"welcome":     "to jurassic park",
		},
	}}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}

func TestMessages_DeleteStatus(t *testing.T) { // TODO

	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"__cursor":"big-cursor","__typename":"Message","archived_at":null,"channel_id":"5da042d7-02ee-46ed-8b91-9b5717da2028","data":{"middle-name":"alfred","welcome":"to jurassic park"},"id":"message-id","inserted_at":"2022-05-17T00:34:18.277163Z","read_at":null,"recipient":"tom","seen_at":null,"source":{"__typename":"NotificationSource","key":"test","version_id":"4dae021a-ba51-473f-9038-77041da8131c"},"status":"delivered","tenant":null,"updated_at":"2022-05-17T00:34:18.318283Z","workflow":"test"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	// ctx, client := RealTestClient() //TODO remove any test client commented code

	have, err := client.Messages.DeleteStatus(ctx, &SetStatusRequest{
		ID:     "29GlZCBxKZ79J9Uquf4HXiUZNfH",
		Status: Archived,
	})

	if err != nil {
		fmt.Println(err)
	}

	want := &SetStatusResponse{Message: &Message{
		Cursor:     "big-cursor",
		ID:         "message-id",
		ChannelID:  "5da042d7-02ee-46ed-8b91-9b5717da2028",
		Recipient:  "tom",
		Workflow:   "test",
		Tenant:     "",
		Status:     "delivered",
		ReadAt:     time.Time{},
		SeenAt:     time.Time{},
		ArchivedAt: time.Time{},
		InsertedAt: time.Date(2022, time.May, 17, 00, 34, 18, 277163000, time.UTC),
		UpdatedAt:  time.Date(2022, time.May, 17, 00, 34, 18, 318283000, time.UTC),
		Source: &NotificationSource{
			Key:       "test",
			VersionID: "4dae021a-ba51-473f-9038-77041da8131c",
		},
		Data: map[string]interface{}{
			"middle-name": "alfred",
			"welcome":     "to jurassic park",
		},
	}}

	c.Assert(err, qt.IsNil)
	c.Assert(have, qt.DeepEquals, want)
}
