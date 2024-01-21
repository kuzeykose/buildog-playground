import { Flex, Text, Button } from "@radix-ui/themes";
import Markdown from "react-markdown";

const md = `# h1 Heading 8-)

## h2 Heading

### h3 Heading

#### h4 Heading

##### h5 Heading

###### h6 Heading
`;

export default function MyApp() {
  return (
    <Flex direction="column" gap="2">
      <Markdown>{md}</Markdown>
    </Flex>
  );
}
