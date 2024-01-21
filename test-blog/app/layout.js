import { Theme } from "@radix-ui/themes";

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <Theme appearance="dark">{children}</Theme>
      </body>
    </html>
  );
}
