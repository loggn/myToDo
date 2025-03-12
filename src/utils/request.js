import axios from 'axios'
// import { useUserStore } from '@/stores'
import { ElMessage } from 'element-plus'

const baseURL = 'http://60.204.240.235:8087'

// 自定义配置新建一个实例
const instance = axios.create({
  baseURL,
  timeout: 10000,
})

// 添加请求拦截器
instance.interceptors.request.use(
  function (config) {
    // 在发送请求之前做些什么
    // 我在这里加入token
    // const useStore = useUserStore()
    // if (useStore.token) {
    //   config.headers.Authorization = useStore.token
    // }

    return config
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error)
  },
)

// 添加响应拦截器
instance.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    if (response.data.code === 0) {
      return response
    }

    return response
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    // 错误的默认情况 => 只要给提示
    ElMessage.error(error.response.data.error || '服务异常')
    return error.response
  },
)

export default instance
export { baseURL }
