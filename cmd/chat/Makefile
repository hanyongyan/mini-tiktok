server:
	kitex -module mini_tiktok idl/chat/chat.proto # execute in the project root directory
	kitex -module mini_tiktok -service chatservice -use mini_tiktok/kitex_gen -I=../../idl/chat chat.proto # execute in cmd/chat
