module.exports = {
  transpileDependencies: true,
  devServer: {
    proxy: 'http://localhost:9090/',
    }
}
