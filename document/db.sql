/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1@3310
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3310
 Source Schema         : cc_main

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 15/06/2020 12:59:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_idmaker
-- ----------------------------
DROP TABLE IF EXISTS `t_idmaker`;
CREATE TABLE `t_idmaker` (
  `id` int(32) NOT NULL AUTO_INCREMENT COMMENT '流水号',
  `key` varchar(64) NOT NULL DEFAULT '' COMMENT 'key',
  `value` int(32) DEFAULT '0' COMMENT 'value',
  `mark` varchar(64) DEFAULT '' COMMENT '备注',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`key`) USING BTREE,
  UNIQUE KEY `key` (`key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COMMENT='Id生成器';

-- ----------------------------
-- Records of t_idmaker
-- ----------------------------
BEGIN;
INSERT INTO `t_idmaker` VALUES (10000, 'uin', 10000, '用户唯一标识', NULL, '2020-05-17 15:31:40', '2020-06-13 09:42:41');
COMMIT;

-- ----------------------------
-- Table structure for t_admin
-- ----------------------------
DROP TABLE IF EXISTS `t_admin`;
CREATE TABLE `t_admin` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `aid` bigint(64) NOT NULL DEFAULT '0' COMMENT '管理员唯一标识',
  `name` varchar(64) DEFAULT '' COMMENT 'admin昵称',
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT 'admin手机号码（账号）',
  `logo` varchar(255) DEFAULT '' COMMENT 'admin头像',
  `sign` varchar(255) DEFAULT '' COMMENT '个性签名',
  `pw` varchar(128) DEFAULT '' COMMENT '密码',
  `salt` varchar(64) DEFAULT '' COMMENT '动态盐',
  `status` tinyint(4) DEFAULT '0' COMMENT 'admin状态（默认0，1在职，2离职）',
  `register_ip` varchar(16) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `modify_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`aid`,`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT COMMENT '流水id',
  `uid` bigint(64) NOT NULL DEFAULT '0' COMMENT '用户唯一标识',
  `last_cid` bigint(64) DEFAULT '0' COMMENT '默认公司id',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `phone` varchar(16) NOT NULL COMMENT '手机号',
  `salt` varchar(32) DEFAULT NULL COMMENT '动态盐',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态(0游客，1公司成员，-1系统拉黑)',
  `loginmode` tinyint(4) DEFAULT NULL COMMENT '登录方式（0有密码，1无密码）',
  `pw` varchar(128) DEFAULT NULL COMMENT '最终密码串',
  `trade_pw` varchar(32) DEFAULT NULL COMMENT '交易密码',
  `idcard` varchar(32) DEFAULT NULL COMMENT '证件号',
  `logo` varchar(255) DEFAULT NULL COMMENT '头像',
  `sign` varchar(255) DEFAULT NULL COMMENT '用户签名',
  `birthday` datetime DEFAULT NULL,
  `sex` tinyint(4) DEFAULT '0' COMMENT '性别（0：男，1：女）',
  `source` varchar(16) DEFAULT NULL COMMENT '来源',
  `register_ip` varchar(16) DEFAULT NULL COMMENT '注册ip',
  `create_time` timestamp NULL DEFAULT NULL,
  `modify_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`uid`,`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for t_company
-- ----------------------------
DROP TABLE IF EXISTS `t_company`;
CREATE TABLE `t_company` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT COMMENT '流水id',
  `cid` bigint(64) NOT NULL COMMENT '公司唯一标识',
  `founder_uid` bigint(64) DEFAULT '0' COMMENT '创始人uid',
  `company` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '公司名称',
  `name` varchar(128) CHARACTER SET utf8 DEFAULT '' COMMENT '法人姓名',
  `idcard` varchar(32) CHARACTER SET utf8 DEFAULT '' COMMENT '法人身份证',
  `phone` varchar(16) CHARACTER SET utf8 DEFAULT '' COMMENT '法人联系电话',
  `logo` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '公司logo',
  `industry` varchar(64) CHARACTER SET utf8 DEFAULT '' COMMENT '公司行业',
  `worktelephone` varchar(64) CHARACTER SET utf8 DEFAULT '' COMMENT '公司单位电话',
  `email` varchar(64) CHARACTER SET utf8 DEFAULT '' COMMENT '公司邮箱地址',
  `website` varchar(64) CHARACTER SET utf8 DEFAULT '' COMMENT '网址',
  `address` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '地址',
  `license_url` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '营业执照（图片路径）',
  `examiner` varchar(255) DEFAULT '' COMMENT '审核人',
  `examine_status` tinyint(4) DEFAULT '0' COMMENT '后台审核状态(0默认：未处理，1通过，-1审核未通过)',
  `remark` text CHARACTER SET utf8 COMMENT '公司简介',
  `applyIp` varchar(16) CHARACTER SET utf8 DEFAULT '' COMMENT '申请来源ip',
  `source` varchar(16) CHARACTER SET utf8 DEFAULT '' COMMENT '来源',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `modify_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`,`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;



-- ----------------------------
-- Table structure for t_message
-- ----------------------------
DROP TABLE IF EXISTS `t_message`;
CREATE TABLE `t_message` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `receive_id` bigint(64) DEFAULT '0' COMMENT '接收人id',
  `send_id` bigint(64) DEFAULT '0' COMMENT '发送者id',
  `logo` varchar(255) DEFAULT NULL COMMENT '系统头像',
  `title` varchar(64) DEFAULT NULL COMMENT '标题',
  `message` text COMMENT '内容',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态（1：已读，0：未读）',
  `send_time` timestamp NULL DEFAULT NULL,
  `type` tinyint(4) DEFAULT NULL COMMENT '消息类型（1：系统，2：公司，3：用户）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_com_industry
-- ----------------------------
DROP TABLE IF EXISTS `t_com_industry`;
CREATE TABLE `t_com_industry` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `industry_type` varchar(64) DEFAULT NULL COMMENT '公司所属行业',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for t_branch
-- ----------------------------
DROP TABLE IF EXISTS `t_branch`;
CREATE TABLE `t_branch` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT COMMENT '流水id',
  `branch_id` bigint(64) NOT NULL DEFAULT '0' COMMENT '部门id',
  `parent_bid` bigint(64) DEFAULT '0' COMMENT '父级部门id',
  `cid` bigint(64) DEFAULT '0' COMMENT '公司id',
  `branch_name` varchar(64) CHARACTER SET utf8 DEFAULT '' COMMENT '部门名称',
  `status` tinyint(4) DEFAULT '1' COMMENT '0默认 1 展示 2 隐藏',
  `create_time` timestamp NULL DEFAULT NULL,
  `modify_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`branch_id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for t_member
-- ----------------------------
DROP TABLE IF EXISTS `t_member`;
CREATE TABLE `t_member` (
  `id` bigint(64) NOT NULL AUTO_INCREMENT,
  `mid` bigint(64) NOT NULL DEFAULT '0' COMMENT '雇员唯一标识',
  `uid` bigint(64) DEFAULT '0' COMMENT '用户id',
  `cid` bigint(64) DEFAULT '0' COMMENT '公司id',
  `bid` bigint(64) DEFAULT '0' COMMENT '部门id',
  `pid` bigint(64) DEFAULT '0' COMMENT '职务id',
  `role` tinyint(4) DEFAULT '3' COMMENT '权限（1：创始人，2：管理员，3：普通员工）',
  `inviter_uid` bigint(64) DEFAULT '0' COMMENT '邀请人uid',
  `work_status` tinyint(4) DEFAULT '1' COMMENT '工作状态（1：工作，-1：离职）',
  `logo` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '员工头像',
  `name` varchar(64) CHARACTER SET utf8 DEFAULT '' COMMENT '员工姓名',
  `idcard` varchar(32) CHARACTER SET utf8 DEFAULT '' COMMENT '身份证号',
  `phone` varchar(16) CHARACTER SET utf8 DEFAULT '' COMMENT '联系电话',
  `sex` tinyint(4) DEFAULT '0' COMMENT '性别（0男，1女）',
  `email` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '员工默认邮箱',
  `create_time` timestamp NULL DEFAULT NULL,
  `modify_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`mid`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=latin1;


-- ----------------------------
-- Table structure for t_form_approval
-- ----------------------------
DROP TABLE IF EXISTS `t_form_approval`;
CREATE TABLE `t_form_approval` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '表单审核流水表主键id',
  `fid` varchar(128) DEFAULT '' COMMENT '提交申请表单时存在的唯一标识符',
  `uid` int(20) DEFAULT '0' COMMENT '审核员id',
  `status` int(20) DEFAULT '0' COMMENT '审核员给出结论（2驳回，1通过）',
  `uuid` varchar(32) DEFAULT '' COMMENT '审核流水记录的唯一标识符',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '流水创建时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_user_form
-- ----------------------------
DROP TABLE IF EXISTS `t_user_form`;
CREATE TABLE `t_user_form` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '用户提交申请任务表单表的主键id',
  `fid` varchar(50) DEFAULT '' COMMENT '申请表单生成单流水号的唯一标识符',
  `uid` int(20) DEFAULT '0' COMMENT '某个用户提交的申请表单',
  `status` int(20) DEFAULT '0' COMMENT '当前表单的审核状态（0草稿，1驳回，2审核中，3审核通过）',
  `cid` int(20) DEFAULT '0' COMMENT '用户在某公司提交的表单对应公司id',
  `submit_way` varchar(50) DEFAULT '' COMMENT '用户的提交方式',
  `source` varchar(50) DEFAULT '' COMMENT '当前表单的来源',
  `result` varchar(512) DEFAULT '' COMMENT '审查结果（唯一标识,多个以''，''拼接）',
  `type` tinyint(2) DEFAULT '0' COMMENT '用户提交表单的类型',
  `content` varchar(512) DEFAULT '' COMMENT '提交申请审核详情',
  `complete_time` datetime DEFAULT NULL COMMENT '单据审核完成时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '当前表单的提交（创建）时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `total_pace` int(11) DEFAULT NULL COMMENT '总进度',
  `currn_pack` int(11) DEFAULT NULL COMMENT '当前进度',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_company_label
-- ----------------------------
DROP TABLE IF EXISTS `t_company_label`;
CREATE TABLE `t_company_label` (
  `id` int(20) NOT NULL AUTO_INCREMENT COMMENT '标签表主键id',
  `name` varchar(50) DEFAULT '' COMMENT '标签名称',
  `integral` int(20) DEFAULT '0' COMMENT '积分增加规则，存放数值代表对应增加多少积分',
  `info` varchar(128) DEFAULT '' COMMENT '标签所需要填写信息的描述',
  `uid` bigint(32) DEFAULT '0' COMMENT '用户id，代表由谁创建的标签',
  `cid` bigint(64) DEFAULT '0' COMMENT '公司cid,代表为那个公司创建',
  `mark` text COMMENT '备注信息',
  `weights` int(11) DEFAULT '0' COMMENT '权重',
  `status` tinyint(2) DEFAULT '0' COMMENT '当前创建标签的状态（0弃用，1草稿，2正式上线）',
  `f_limit` int(20) DEFAULT '0' COMMENT '表示当前标签一天每个用户最多能提交多少次',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8mb4;