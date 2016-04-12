import React from 'react';
import CodeSlide from 'spectacle-code-slide';

// Import Spectacle Core tags
import {
  Appear,
  BlockQuote,
  Cite,
  CodePane,
  Deck,
  Fill,
  Heading,
  Image,
  Layout,
  Link,
  ListItem,
  List,
  Markdown,
  Quote,
  Slide,
  Spectacle,
  Text
} from "spectacle";

// Import theme
import createTheme from "spectacle/lib/themes/default";

// Require CSS
require("normalize.css");
require("spectacle/lib/themes/default/index.css");

const theme = createTheme({
  primary: "#ff4081"
});

export default class Presentation extends React.Component {
  render() {
    return (
      <Spectacle theme={theme}>
        <Deck transition={[]} transitionDuration={0} progress="bar">
          <Slide transition={["zoom"]} bgColor="primary">
            <Heading size={1} fit caps lineHeight={1} textColor="black">
              Building Websockets with Go
            </Heading>
            <Heading size={1} fit caps>
              Code Walk Through
            </Heading>
          </Slide>
          <CodeSlide
            transition={[]}
            lang="js"
            code={require("raw!./assets/server.example")}
            ranges={[
              { loc: [0, 0], title: "Server Code" },
              { loc: [0, 9]},
              { loc: [11, 21] },
              { loc: [23, 35] },
              { loc: [50, 63] },
              { loc: [56, 57], note: "This is where the connection is added"},
              { loc: [66, 75] },
              { loc: [75, 90] },
              { loc: [37, 42] },
            ]}/>
          <CodeSlide
            transition={[]}
            lang="js"
            code={require("raw!./assets/connection.example")}
            ranges={[
              { loc: [0, 0], title: "Connection Code" },
              { loc: [24, 31]},
              { loc: [33, 46] },
              { loc: [16, 17] },
              { loc: [57, 61] },
              { loc: [81, 82], note: "Receiving messages"},
              { loc: [82, 90] },
              { loc: [92, 109] },
              { loc: [11, 15], note: "JSON format" },
              { loc: [63, 79], note: "Sending messages"},
            ]}/>
          <Slide>
            <Heading size={1} fit caps lineHeight={1} textColor="black">
              The End
            </Heading>
          </Slide>
        </Deck>
      </Spectacle>
    );
  }
}
