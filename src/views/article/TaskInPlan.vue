<template>
  <div class="TaskInPlan-common-layout">
    <el-container>
      <el-header class="TaskInPlan-header">
        <div class="TaskInPlan-header-left">
          <h1>
            <el-icon><List /></el-icon>
            计划内
          </h1>
        </div>
        <div class="TaskInPlan-header-rigth">
          <div>
            <el-icon><TopRight /></el-icon>
          </div>
          <div>...</div>
        </div>
      </el-header>
      <el-main class="TaskInPlan-body">任务列表</el-main>
      <el-footer class="TaskInPlan-foot">添加任务</el-footer>
    </el-container>
  </div>
</template>
<style>
.TaskInPlan-common-layout {
  display: flex;
  background-image: url('../../assets/背景4.jpg');
  background-size: 100% 100%;
  background-repeat: no-repeat;
  height: 100vh;
}
.TaskInPlan-header {
  display: flex;
  justify-content: space-between;
  margin: 35px;
  height: 135px;
}
.TaskInPlan-header-left {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: white;
}

.TaskInPlan-header-rigth {
  display: flex;
  justify-content: space-around;
  align-items: center;
  font-size: 25px;
  color: white;
  width: 180px;
}
</style>
