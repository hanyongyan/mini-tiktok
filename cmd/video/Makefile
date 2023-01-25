server:
	kitex -module mini_tiktok idl/video/video.proto # execute in the project root directory
	kitex -module mini_tiktok -service videoservice -use mini_tiktok/kitex_gen -I=../../idl/video video.proto # execute in cmd/video
