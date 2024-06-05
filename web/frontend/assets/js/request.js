; (function (win) {
  // 创建axios实例
  const service = axios.create({
    // axios中请求配置有baseURL选项，表示请求URL公共部分
    baseURL: '/',
    // 超时
    timeout: 10000,
    // 请求头
    headers: { 'Content-Type': 'application/json;charset=utf-8' },
  })

  service.defaults.retry = 3 // 请求重试次数
  service.defaults.retryDelay = 1000 // 请求重试时间间隔
  service.defaults.shouldRetry = true // 是否重试

  // request拦截器
  service.interceptors.request.use(
    (config) => {
      //添加header
      const token = localStorage.getItem('token')
      if (token) {
        // 添加请求头
        config.headers['Authorization'] = 'Bearer ' + token
      }
      // get请求映射params参数
      if (config.method === 'get' && config.params) {
        let url = config.url + '?'
        for (const propName of Object.keys(config.params)) {
          const value = config.params[propName]
          var part = encodeURIComponent(propName) + '='
          if (value !== null && typeof value !== 'undefined') {
            if (typeof value === 'object') {
              for (const key of Object.keys(value)) {
                let params = propName + '[' + key + ']'
                var subPart = encodeURIComponent(params) + '='
                url += subPart + encodeURIComponent(value[key]) + '&'
              }
            } else {
              url += part + encodeURIComponent(value) + '&'
            }
          }
        }
        url = url.slice(0, -1)
        config.params = {}
        config.url = url
      }
      return config
    },
    (error) => {
      console.log(error)
      Promise.reject(error)
    }
  )

  // 响应拦截器
  service.interceptors.response.use(
    (response) => {
      const { code, msg } = response.data
      if (code === 0 && msg === 'NOTLOGIN') {
        // 返回登录页面
        localStorage.removeItem('userInfo')
        window.top.location.href = '/front/page/login/login.html'
      }
      if (code === 200) {
        return response.data
      }
      // 响应数据为二进制流处理(Excel导出)
      if (response.data instanceof ArrayBuffer) {
        return response
      }
    },
    (error) => {
      console.log('err' + error)
      let { message } = error
      if (message == 'Network Error') {
        message = '后端接口连接异常'
      } else if (message.includes('timeout')) {
        message = '系统接口请求超时'
      } else if (message.includes('Request failed with status code')) {
        message = '系统接口' + message.substr(message.length - 3) + '异常'
      }
      // 注意，这里使用ElementPlusUI的ElMessage组件
      window.ElementPlus.ElMessage({
        message: message,
        type: 'error',
        duration: 5 * 1000,
      })
      return Promise.reject(error)
    }
  )
  win.$axios = service
})(window)
