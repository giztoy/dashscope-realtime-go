package dashscope

import "testing"

func TestDecodeServerEventChoicesFormat(t *testing.T) {
	input := []byte(`{
		"choices": [
			{
				"finish_reason": "stop",
				"message": {
					"content": [
						{"text": "hello"},
						{"audio": {"data": "aGVsbG8="}}
					]
				}
			}
		]
	}`)

	event, err := DecodeServerEvent(input)
	if err != nil {
		t.Fatalf("DecodeServerEvent() error = %v", err)
	}

	if event.Type != "choices" {
		t.Fatalf("event.Type = %q, want %q", event.Type, "choices")
	}
	if event.Delta != "hello" {
		t.Fatalf("event.Delta = %q, want %q", event.Delta, "hello")
	}
	if event.AudioBase64 != "aGVsbG8=" {
		t.Fatalf("event.AudioBase64 = %q, want %q", event.AudioBase64, "aGVsbG8=")
	}
	if event.FinishReason != "stop" {
		t.Fatalf("event.FinishReason = %q, want %q", event.FinishReason, "stop")
	}
}

func TestDecodeServerEventError(t *testing.T) {
	input := []byte(`{
		"type": "error",
		"error": {
			"type": "invalid_request_error",
			"code": "InvalidApiKey",
			"message": "invalid token"
		}
	}`)

	event, err := DecodeServerEvent(input)
	if err != nil {
		t.Fatalf("DecodeServerEvent() error = %v", err)
	}

	if event.Type != "error" {
		t.Fatalf("event.Type = %q, want %q", event.Type, "error")
	}
	if event.Error == nil {
		t.Fatalf("event.Error is nil")
	}
	if event.Error.Code != "InvalidApiKey" {
		t.Fatalf("event.Error.Code = %q, want %q", event.Error.Code, "InvalidApiKey")
	}
}

func TestDecodeServerEventUsage(t *testing.T) {
	input := []byte(`{
		"type": "response.done",
		"response": {
			"id": "resp_1",
			"usage": {
				"total_tokens": 3,
				"input_tokens": 1,
				"output_tokens": 2
			}
		}
	}`)

	event, err := DecodeServerEvent(input)
	if err != nil {
		t.Fatalf("DecodeServerEvent() error = %v", err)
	}
	if event.ResponseID != "resp_1" {
		t.Fatalf("event.ResponseID = %q, want %q", event.ResponseID, "resp_1")
	}
	if event.Usage == nil {
		t.Fatalf("event.Usage is nil")
	}
	if event.Usage.TotalTokens != 3 || event.Usage.InputTokens != 1 || event.Usage.OutputTokens != 2 {
		t.Fatalf("unexpected usage: %#v", event.Usage)
	}
}

func TestDecodeServerEventOutputItemFields(t *testing.T) {
	input := []byte(`{
		"type": "response.output_item.added",
		"response_id": "resp_2",
		"output_index": 1,
		"item": {
			"id": "item_9",
			"type": "message",
			"role": "assistant",
			"status": "in_progress",
			"content": [
				{"type": "text", "text": "hello"}
			]
		}
	}`)

	event, err := DecodeServerEvent(input)
	if err != nil {
		t.Fatalf("DecodeServerEvent() error = %v", err)
	}

	if event.Type != "response.output_item.added" {
		t.Fatalf("event.Type = %q, want %q", event.Type, "response.output_item.added")
	}
	if event.ResponseID != "resp_2" {
		t.Fatalf("event.ResponseID = %q, want %q", event.ResponseID, "resp_2")
	}
	if event.OutputIndex != 1 {
		t.Fatalf("event.OutputIndex = %d, want 1", event.OutputIndex)
	}
	if event.ItemID != "item_9" {
		t.Fatalf("event.ItemID = %q, want %q", event.ItemID, "item_9")
	}
	if event.Item == nil {
		t.Fatalf("event.Item is nil")
	}
	if got := len(event.Item.Content); got != 1 {
		t.Fatalf("len(event.Item.Content) = %d, want 1", got)
	}
	if event.Delta != "hello" {
		t.Fatalf("event.Delta = %q, want %q", event.Delta, "hello")
	}
}

func TestDecodeServerEventResponseField(t *testing.T) {
	input := []byte(`{
		"type": "response.created",
		"response": {
			"id": "resp_3",
			"status": "in_progress",
			"output": [{
				"id": "item_1",
				"type": "message",
				"role": "assistant",
				"status": "in_progress",
				"content": [{"type":"text","text":"hi"}]
			}]
		}
	}`)

	event, err := DecodeServerEvent(input)
	if err != nil {
		t.Fatalf("DecodeServerEvent() error = %v", err)
	}

	if event.Response == nil {
		t.Fatalf("event.Response is nil")
	}
	if event.Response.ID != "resp_3" {
		t.Fatalf("event.Response.ID = %q, want %q", event.Response.ID, "resp_3")
	}
	if event.ResponseID != "resp_3" {
		t.Fatalf("event.ResponseID = %q, want %q", event.ResponseID, "resp_3")
	}
	if got := len(event.Response.Output); got != 1 {
		t.Fatalf("len(event.Response.Output) = %d, want 1", got)
	}
}
