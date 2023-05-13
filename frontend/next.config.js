/** @type {import('next').NextConfig} */
const nextConfig = {
    experimental: {
        appDir: true,
    },
    output: "export",
    distDir: '../backend/views',
    // Cloudflareのパブリック r2.dev バケット URL
    assetPrefix: process.env.STATIC_FILE_URL,
}

module.exports = nextConfig
