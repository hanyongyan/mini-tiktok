server:
	kitex -module mini_tiktok idl/user/user.proto # execute in the project root directory
	kitex -module mini_kitex -module minitiktok -service userservice -use mini_tiktok/kitex_gen -I=../../idl/user user.proto # execute in cmd/user
