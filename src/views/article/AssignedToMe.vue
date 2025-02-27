<template>
  <div class="AssignedToMe-common-layout">
    <el-container>
      <el-header class="AssignedToMe-header">
        <div class="AssignedToMe-header-left">
          <h1>
            <el-icon><User /></el-icon>
            分配给我的
          </h1>
        </div>
        <div class="AssignedToMe-header-rigth">
          <div>
            <el-icon><TopRight /></el-icon>
          </div>
          <div>...</div>
        </div>
      </el-header>
      <el-main class="AssignedToMe-body">任务列表</el-main>
      <el-footer class="AssignedToMe-foot">添加任务</el-footer>
    </el-container>
  </div>
</template>
<style>
.AssignedToMe-common-layout {
  display: flex;
  background-image: url('../../assets/背景1.jpg');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  height: 100vh;
}
.AssignedToMe-header {
  display: flex;
  justify-content: space-between;
  margin: 35px;
  height: 135px;
}
.AssignedToMe-header-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: white;
}

.AssignedToMe-header-rigth {
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 25px;
  color: white;
  width: 180px;
}
</style>
