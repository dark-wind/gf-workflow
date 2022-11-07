
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="../public/front.css">

    <title>Document</title>
</head>
<body>
<div id="main">
    <div id="left">
        <h6>taskId：<span id="taskId">{{.taskId}}</span></h6>
        <h6>发起用户名称：{{.StartUserName}}</h6>
        <h6>当前节点名称：{{.NodeName}}</h6>
        <h6>流程名称：{{.ProcessName}}</h6>
        <h6>应当处理该任务的角色：{{.AssigneeRoleName}}</h6>
        <h6>任务的条件值：{{.Conditions}}</h6>
        <h6>流程状态：{{.Status}}</h6>
    </div>
    <div id="right">
        <div id="main1">
            <div class="area">
                <div>
                    <h6>用户名：菜虚鲲</h6>
                    <br>
                    <label for="Status">政治面貌：</label>
                    <select name="" id="Status"><br>
                        <option value="党员">党员</option>
                        <option value="非党员">非党员</option>
                    </select>
                    <br>
                    <label for="Type">流程选择：</label>
                    <select name="" id="Type"><br>
                        {{range $key, $value := .list}}
                        <option value="{{$key}}">{{$value}}</option>
                        {{end}}
                    </select>
                    <button id="start-button">发起流程</button>
                </div>

            </div>
            <div class="area">
                <h6>用户名：路飞</h6>
                <h6>用户角色：招生办</h6>
                <table border="1">
                    <tr>
                        <th>发起人姓名</th>
                        <th>审批类型</th>
                        <th>操作</th>
                    </tr>
                    {{if .zsbStartUserName}}
                    <tr>
                        <th>{{.zsbStartUserName}}</th>
                        <th>{{.zsbProcessName}}</th>
                        <th><button id="zsb-submit">通过</button>
                            <button onclick="refuse">驳回</button>
                        </th>
                    </tr>
                    {{end}}
                </table>
            </div>
            <div class="area">
                <h6>用户名：冯宝宝</h6>
                <h6>用户角色：辅导员</h6>
                <table border="1">
                    <tr>
                        <th>发起人姓名</th>
                        <th>审批类型</th>
                        <th>操作</th>
                    </tr>
                    {{if .fdyStartUserName}}
                    <tr>
                        <th>{{.fdyStartUserName}}</th>
                        <th>{{.fdyProcessName}}</th>
                        <th><button id="fdy-submit" onclick="pass()">通过</button>
                            <button onclick="refuse">驳回</button>
                        </th>
                    </tr>
                    {{end}}
                </table>
            </div>
        </div>
        <div id="main2">
            <div class="area">
                <h6>用户名：范闲</h6>
                <h6>用户角色：财务处</h6>
                <table border="1">
                    <tr>
                        <th>发起人姓名</th>
                        <th>审批类型</th>
                        <th>操作</th>
                    </tr>
                    {{if .cwcStartUserName}}
                    <tr>
                        <th>{{.cwcStartUserName}}</th>
                        <th>{{.cwcProcessName}}</th>
                        <th><button id="cwc-submit">通过</button>
                            <button onclick="refuse">驳回</button>
                        </th>
                    </tr>
                    {{end}}
                </table>
            </div>
            <div class="area">
                <h6>用户名：阿贝尔</h6>
                <h6>用户角色：导师</h6>
                <table border="1">
                    <tr>
                        <th>发起人姓名</th>
                        <th>审批类型</th>
                        <th>操作</th>
                    </tr>
                    {{if .dsStartUserName}}
                    <tr>
                        <th>{{.dsStartUserName}}</th>
                        <th>{{.dsProcessName}}</th>
                        <th><button id="ds-submit">通过</button>
                            <button onclick="refuse">驳回</button>
                        </th>
                    </tr>
                    {{end}}
                </table>
            </div>
            <div class="area">
                <h6>用户名：五条悟</h6>
                <h6>用户角色：党支部</h6>
                <table border="1">
                    <tr>
                        <th>发起人姓名</th>
                        <th>审批类型</th>
                        <th>操作</th>
                    </tr>
                    {{if .dzbStartUserName}}
                    <tr>
                        <th>{{.dzbStartUserName}}</th>
                        <th>{{.dzbProcessName}}</th>
                        <th><button id="dzb-submit">通过</button>
                            <button onclick="refuse">驳回</button>
                        </th>
                    </tr>
                    {{end}}
                </table>
            </div>

        </div>
        <div id="main3">
            <div class="area">
                <h6>用户名：川建国</h6>
                <h6>用户角色：团支部</h6>
                <table border="1">
                    <tr>
                        <th>发起人姓名</th>
                        <th>审批类型</th>
                        <th>操作</th>
                    </tr>
                    {{if .tzbStartUserName}}
                    <tr>
                        <th>{{.tzbStartUserName}}</th>
                        <th>{{.tzbProcessName}}</th>
                        <th><button id="tzb-submit">通过</button>
                            <button onclick="refuse">驳回</button>
                        </th>
                    </tr>
                    {{end}}
                </table>
            </div>
            <div class="area">
                <h6>用户名：栗子姨</h6>
                <h6>用户角色：宿管</h6>
                <table border="1">
                    <tr>
                        <th>发起人姓名</th>
                        <th>审批类型</th>
                        <th>操作</th>
                    </tr>
                    {{if .sgStartUserName}}
                    <tr>
                        <th>{{.sgStartUserName}}</th>
                        <th>{{.sgProcessName}}</th>
                        <th><button id="sg-submit">通过</button>
                            <button onclick="refuse">驳回</button>
                        </th>
                    </tr>
                    {{end}}
                </table>
            </div>
        </div>
    </div>

</div>
</div>
<script src="../public/jquery-3.6.1.min.js"></script>
<script src="../public/front.js"></script>
</body>
</html>