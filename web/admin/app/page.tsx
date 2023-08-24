// "use client";
import VideoPlayer from "@/components/videoplayer";
import { Button } from "@radix-ui/themes";

export default function Home() {
  return (
    <>
      <Button>Hello</Button>
      <VideoPlayer src={"http://localhost:4000/index.m3u8"} />
    </>
  );
}
