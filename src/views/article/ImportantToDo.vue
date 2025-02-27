<template>
  <div class="Import-common-layout">
    <el-container>
      <el-header class="Import-header">
        <div class="Import-header-left">
          <h1>
            <el-icon><StarFilled /></el-icon>
            重要
          </h1>
        </div>
        <div class="Import-header-rigth">
          <div>
            <el-icon><TopRight /></el-icon>
          </div>
          <div>...</div>
        </div>
      </el-header>
      <el-main class="Import-body">任务列表</el-main>
      <el-footer class="Import-foot">添加任务</el-footer>
    </el-container>
  </div>
</template>
<style>
.Import-common-layout {
  display: flex;
  background-image: url('../../assets/背景3.jpg');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  height: 100vh;
}
.Import-header {
  display: flex;
  justify-content: space-between;
  margin: 35px;
  height: 135px;
}
.Import-header-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: white;
}

.Import-header-rigth {
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 25px;
  color: white;
  width: 180px;
}
</style>
