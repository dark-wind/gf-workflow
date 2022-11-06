<!DOCTYPE html>
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
    <div id="main1">
        <div class="area">
            <form action="">
                <label for="UserName">用户名：</label>
                <input type="text" id="UserName" align="left">
                <br>
                <label for="Status">政治面貌：</label>
                <select name="" id="Status"><br>
                    <option>党员</option>
                    <option>共青团员</option>
                </select>
                <br>
                <label for="Type">申请类型：</label>
                <select name="" id="Type"><br>
                    <option>报道</option>
                </select>
                <input type="submit" value="提交">
            </form>

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
                <tr>
                    <th>{{.zsbStartUserName}}</th>
                    <th>{{.zsbProcessName}}</th>
                    <th><button>通过</button>
                        <button>驳回</button>
                    </th>
                </tr>
            </table>
        </div>
        <div class="area">
            <h6>用户名：路飞</h6>
            <h6>用户角色：辅导员</h6>
            <table border="1">
                <tr>
                    <th>发起人姓名</th>
                    <th>审批类型</th>
                    <th>操作</th>
                </tr>
                <tr>
                    <th>{{.fdyStartUserName}}</th>
                    <th>{{.fdyProcessName}}</th>
                    <th><button>通过</button>
                        <button>驳回</button>
                    </th>
                </tr>
            </table>
        </div>
    </div>
    <div id="main2">
        <div class="area">
            <h6>用户名：路飞</h6>
            <h6>用户角色：财务处</h6>
            <table border="1">
                <tr>
                    <th>发起人姓名</th>
                    <th>审批类型</th>
                    <th>操作</th>
                </tr>
                <tr>
                    <th>{{.cwcStartUserName}}</th>
                    <th>{{.cwcProcessName}}</th>
                    <th><button>通过</button>
                        <button>驳回</button>
                    </th>
                </tr>
            </table>
        </div>
        <div class="area">
            <h6>用户名：路飞</h6>
            <h6>用户角色：导师</h6>
            <table border="1">
                <tr>
                    <th>发起人姓名</th>
                    <th>审批类型</th>
                    <th>操作</th>
                </tr>
                <tr>
                    <th>{{.dsStartUserName}}</th>
                    <th>{{.dsProcessName}}</th>
                    <th><button>通过</button>
                        <button>驳回</button>
                    </th>
                </tr>
            </table>
        </div>
        <div class="area">
            <h6>用户名：路飞</h6>
            <h6>用户角色：党支部</h6>
            <table border="1">
                <tr>
                    <th>发起人姓名</th>
                    <th>审批类型</th>
                    <th>操作</th>
                </tr>
                <tr>
                    <th>{{.dzbStartUserName}}</th>
                    <th>{{.dzbProcessName}}</th>
                    <th><button>通过</button>
                        <button>驳回</button>
                    </th>
                </tr>
            </table>
        </div>

    </div>
    <div id="main3">
        <div class="area">
            <h6>用户名：路飞</h6>
            <h6>用户角色：团支部</h6>
            <table border="1">
                <tr>
                    <th>发起人姓名</th>
                    <th>审批类型</th>
                    <th>操作</th>
                </tr>
                <tr>
                    <th>{{.tzbStartUserName}}</th>
                    <th>{{.tzbProcessName}}</th>
                    <th><button>通过</button>
                        <button>驳回</button>
                    </th>
                </tr>
            </table>
        </div>
        <div class="area">
            <h6>用户名：路飞</h6>
            <h6>用户角色：宿管</h6>
            <table border="1">
                <tr>
                    <th>发起人姓名</th>
                    <th>审批类型</th>
                    <th>操作</th>
                </tr>
                <tr>
                    <th>{{.sgStartUserName}}</th>
                    <th>{{.sgProcessName}}</th>
                    <th><button>通过</button>
                        <button>驳回</button>
                    </th>
                </tr>
            </table>
        </div>

    </div>

</div>
</div>
</body>
</html>