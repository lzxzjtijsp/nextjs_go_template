/** @type {import('next').NextConfig} */
const nextConfig = {
    experimental: {
        appDir: true,
    },
    output: "export",
    distDir: '/views',
    // Cloudflareのパブリック r2.dev バケット URL
    assetPrefix: process.env.BUCKET_URL,
}

module.exports = nextConfig
