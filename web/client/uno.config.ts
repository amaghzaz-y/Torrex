import {
  defineConfig,
  presetAttributify,
  presetIcons,
  presetTypography,
  presetUno,
  presetWebFonts,
  transformerDirectives,
  transformerVariantGroup,
} from "unocss";

export default defineConfig({
  shortcuts: [
    // ...
  ],
  theme: {
    colors: {
      torrex: {
        primary: "rgba(136, 6, 37, 20%)",
        secondary: "rbga(190,190,190,5%)",
        accent: "#CF0031",
        background: "rgba(15,0,5,100%)",
        text: "rgba(205,220,220,80%)",
        paragraph: "rgba(255,0,60,80%)",
      },
    },
  },
  presets: [
    presetUno(),
    presetAttributify(),
    presetIcons(),
    presetTypography(),
    presetWebFonts({
      provider: "google",
      fonts: {
        sans: [
          "Poppins",
          "Poppins:400",
          "Poppins:500",
          "Poppins:600",
          "Poppins:700",
          "Poppins:800",
          "Poppins:900",
        ],
      },
    }),
  ],
  transformers: [transformerDirectives(), transformerVariantGroup()],
});
