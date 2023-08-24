import { Movie } from "@/lib/movie";
import { Card, Flex, Heading, Text } from "@radix-ui/themes";

export default function MovieCard({ movie }: { movie: Movie }) {
  return (
    <>
      <Card>
        <Flex>
          <Heading>{movie.title}</Heading>
          <Text>{movie.tagline}</Text>
          <Text>{movie.description}</Text>
          <Text>{movie.score}</Text>
          <Text>{movie.year}</Text>
        </Flex>
      </Card>
    </>
  );
}
