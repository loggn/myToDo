import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore(
  'big-user',
  () => {
    const token = ref('')
    const user_id = ref('')
    const name = ref('')
    const setToken = (newToken) => {
      token.value = newToken
    }
    const removeToken = () => {
      token.value = ''
    }
    const setUserId = (newUserId) => {
      user_id.value = newUserId
    }
    const removeUserId = () => {
      user_id.value = ''
    }
    const setName = (newName) => {
      name.value = newName
    }
    const removeName = () => {
      name.value = ''
    }

    return {
      user_id,
      name,
      token,
      setToken,
      removeToken,
      setUserId,
      removeUserId,
      setName,
      removeName,
    }
  },
  {
    persist: true,
  },
)
