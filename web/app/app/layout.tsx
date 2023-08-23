import "@radix-ui/themes/styles.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { Theme } from "@radix-ui/themes";
const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Pied Piper",
  description: "The Robinhood of Hollywood",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <Theme>
        <body className={inter.className}>{children}</body>
      </Theme>
    </html>
  );
}
