/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ["class"],
  content: ["./pages/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}", "./app/**/*.{ts,tsx}", "./src/**/*.{ts,tsx}"],
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      colors: {
        // ☆少し濃い
        header: {
          DEFAULT: "#FFF9E6", // ヘッダーの背景色（ほぼ白に近い薄オレンジ）
          text: "#829ba1", // ヘッダーのテキスト色（柔らかいダークグレー）
        },
        sidebar: {
          DEFAULT: "#F7FFFF", // サイドバーの背景色（ほぼ白に近い薄水色）
          text: "#829ba1", // サイドバーのテキスト色（読みやすいグレー）
        },
        chat: {
          DEFAULT: "#FFFFFF", // チャットの背景色（純白）
          text: "#829ba1", // チャットのテキスト色（濃いグレー）
          accent: "#FFFEF7", // チャットのアクセント色（ほぼ白に近い薄オレンジ）
        },
        online: {
          DEFAULT: "#F7FFFF", // オンライン/オフラインバーの背景色（ほぼ白に近い薄水色）
          text: "#829ba1", // オンライン/オフラインバーのテキスト色（サイドバーと同じ色）
        },

        // header: {
        //   DEFAULT: "#FFF7E0", // ヘッダーの背景色（薄黄色、より控えめ）
        //   text: "#424242", // ヘッダーのテキスト色（柔らかいダークグレー）
        // },
        // sidebar: {
        //   DEFAULT: "#E0F7FA", // サイドバーの背景色（薄水色、より控えめ）
        //   text: "#616161", // サイドバーのテキスト色（読みやすいグレー）
        // },
        // chat: {
        //   DEFAULT: "#FFFFFF", // チャットの背景色（純白）
        //   text: "#212121", // チャットのテキスト色（濃いグレー）
        //   accent: "#FFF7E0", // チャットのアクセント色（薄黄色）
        // },
        // online: {
        //   DEFAULT: "#E0F7FA", // オンライン/オフラインバーの背景色（薄水色）
        //   text: "#616161", // オンライン/オフラインバーのテキスト色（サイドバーと同じ色）
        // },

        //ライトモード
        // header: {
        //   DEFAULT: "#F0F8FF", // ヘッダーの背景色（非常に淡い水色）
        //   text: "#4A4A4A", // ヘッダーのテキスト色（やや濃いグレー）
        // },
        // sidebar: {
        //   DEFAULT: "#F0F8FF", // サイドバーの背景色（非常に淡い水色）
        //   text: "#4A4A4A", // サイドバーのテキスト色（やや濃いグレー）
        // },
        // chat: {
        //   DEFAULT: "#F0F8FF", // チャットの背景色（非常に淡い水色）
        //   text: "#4A4A4A", // チャットのテキスト色（やや濃いグレー）
        // },
        // online: {
        //   DEFAULT: "#F0F8FF", // オンライン/オフラインバーの背景色（非常に淡い水色）
        //   text: "#4A4A4A", // オンライン/オフラインバーのテキスト色（やや濃いグレー）
        // },

        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsl(var(--background))",
        foreground: "hsl(var(--foreground))",
        primary: {
          DEFAULT: "hsl(var(--primary))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsl(var(--secondary))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsl(var(--destructive))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsl(var(--muted))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsl(var(--accent))",
          foreground: "hsl(var(--accent-foreground))",
        },
        popover: {
          DEFAULT: "hsl(var(--popover))",
          foreground: "hsl(var(--popover-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card))",
          foreground: "hsl(var(--card-foreground))",
        },
      },
      borderRadius: {
        lg: "var(--radius)",
        md: "calc(var(--radius) - 2px)",
        sm: "calc(var(--radius) - 4px)",
      },
      keyframes: {
        "accordion-down": {
          from: { height: 0 },
          to: { height: "var(--radix-accordion-content-height)" },
        },
        "accordion-up": {
          from: { height: "var(--radix-accordion-content-height)" },
          to: { height: 0 },
        },
      },
      animation: {
        "accordion-down": "accordion-down 0.2s ease-out",
        "accordion-up": "accordion-up 0.2s ease-out",
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
};
