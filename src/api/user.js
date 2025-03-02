import request from '@/utils/request'

export const userRegisterService = (account, password, repassword) => {
  return request.post('/API/register', {
    account: account,
    password: password,
    repassword: repassword
  })
}

export const userLoginService = (account, password) => {
  return request.post('/API/login', {
    account: account,
    password: password,
  })
}

export const userAddTask = (text, user_id) => {
  return request.post('/API/addTask', {
    user_id: user_id,
    text: text,
  })
}

export const userGetTask = (user_id) => {
  return request.post('/API/getTask', {
    id: user_id,
  })
}

export const changeUserName = (user_id, name) => {
  return request.post('/API/changeUserName', {
    id: user_id,
    name: name,
  })
}

export const changeTaskText = (task_id, task_text) => {
  return request.post('/API/changeTaskText', {
    id: task_id,
    text: task_text,
  })
}

export const changeTaskIsFinish = (task_id, task_isFinish) => {
  return request.post('/API/changeTaskIsFinish', {
    id: task_id,
    isFinish: task_isFinish,
  })
}

export const changeTaskIsImportant = (task_id, task_isImportant) => {
  return request.post('/API/changeTaskIsImportant', {
    id: task_id,
    isImportant: task_isImportant,
  })
}

export const deleteTask = (task_id) => {
  return request.post('/API/deleteTask', {
    id: task_id,
  })
}
