"use client";
import { MagnifyingGlassIcon } from "@radix-ui/react-icons";
import { Button, Container, Flex, Section, TextField } from "@radix-ui/themes";

export default function Home() {
  return (
    <Section>
      <Flex justify={"center"} align={"center"} gap={"4"}>
        <TextField.Root style={{ width: "30rem" }}>
          <TextField.Slot>
            <MagnifyingGlassIcon height="24" width="24" />
          </TextField.Slot>
          <TextField.Input placeholder="Search Movies" size={"2"} />
        </TextField.Root>
        <Button size={"2"}>Search</Button>
      </Flex>
    </Section>
  );
}
