package viewmodel

import (
	"backend/internal/model"
)

type BaseChatView struct {
	TelegramChatID int64 `json:"telegram_chat_id"`
}

type ChatListView struct {
	Items []ChatView `json:"items"`
}
type ChatView struct {
	BaseChatView
	RequestQueries []RequestQueryView `json:"requestQueries"`
}

func ChatListViewFrom(chats []model.Chat) ChatListView {
	chatViews := make([]ChatView, len(chats))

	for i, chat := range chats {
		chatViews[i] = ChatViewFrom(&chat)
	}
	return ChatListView{
		Items: chatViews,
	}
}

func ChatViewFrom(c *model.Chat) ChatView {
	requestQueryViews := make([]RequestQueryView, len(c.RequestQueries))
	for i, rq := range c.RequestQueries {
		requestQueryViews[i] = RequestQueryViewFrom(&rq)
	}
	return ChatView{
		BaseChatView: BaseChatView{
			TelegramChatID: c.TelegramChatId,
		},
		RequestQueries: requestQueryViews,
	}
}
