<svelte:head>
  <!-- Fonts --><link
    rel="stylesheet"
    href="https://fonts.googleapis.com/icon?family=Material+Icons"
  />
  <link
    rel="stylesheet"
    href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,600,700"
  />

  <!-- Material Typography --><link
    rel="stylesheet"
    href="https://unpkg.com/@material/typography@13.0.0/dist/mdc.typography.css"
  />

  <!-- SMUI --><link rel="stylesheet" href="https://unpkg.com/svelte-material-ui/bare.css" />
</svelte:head>

<script>
	import Button from '@smui/button';
  import Card, { Content } from '@smui/card';
	import { Router, Link, Route } from "svelte-routing";
  import LogoutButton from './Login/LogoutButton.svelte';
  import MfaSetupButton from './Login/MfaSetupButton.svelte';
	import UserGroupManager from "./UserGroup/UserGroupManager.svelte";
	import AddUserGroup from "./UserGroup/AddUserGroup.svelte";
	import EditUserGroup from "./UserGroup/UserGroupEdit.svelte";
  import UserManager from "./User/UserManager.svelte";
  import AddUser from './User/AddUser.svelte';
  import EditUser from "./User/UserEdit.svelte";
  import EditUserDevice from "./UserDevice/UserDeviceEdit.svelte";
  import AddUserDevice from "./UserDevice/AddUserDevice.svelte";
  import Register from "./Login/Register.svelte";
  import Login from "./Login/Login.svelte";
  import ConfirmAccount from "./Login/ConfirmAccount.svelte";
  import MfaSetup from "./Login/MfaSetup.svelte";
  import Logout from "./Login/Logout.svelte";
  import DeleteUserPage from "./Login/DeleteUserPage.svelte";
  import Enter from "./Enter.svelte";
  import UserDashboard from "./Dashboards/UserDashboard.svelte";
  import UserGroupDetail from "./UserGroup/UserGroupDetail.svelte";
  import AddUserToGroup from "./UserGroup/AddUserToGroup.svelte";
  import UserDetail from "./User/UserDetail.svelte";
  import UserDeviceDetail from "./UserDevice/UserDeviceDetail.svelte";
  import AdminDashboard from './Dashboards/AdminDashboard.svelte';

	export let url = "";
</script>

<Router url="{url}">
	<div>
		<Route path="/" component="{Enter}" />
    <Route path="usergroup" component="{UserGroupManager}" />
		<Route path="usergroup/add" component="{AddUserGroup}" />
		<Route path="usergroup/edit/:id" component="{EditUserGroup}" />
    <Route path="usergroup/detail/:id" component="{UserGroupDetail}" />
    <Route path="usergroup/detail/:id/adduser" component="{AddUserToGroup}" />
    <Route path="user" component="{UserManager}" />
    <Route path="user/add" component="{AddUser}" />
		<Route path="user/edit/:id" component="{EditUser}" />
    <Route path="user/detail/:id" component="{UserDetail}" />
    <Route path="user/detail/:id/adddevice" component="{AddUserDevice}" />
    <Route path="user/detail/:id/editdevice/:did" component="{EditUserDevice}" />
    <Route path="user/detail/:id/device/:did" component="{UserDeviceDetail}" />
    <Route path="login" component="{Login}" />
    <Route path="confirm" component="{ConfirmAccount}" />
    <Route path="mfa" component="{MfaSetup}" />
    <Route path="register" component="{Register}" />
    <Route path="logout" component="{Logout}" />
    <Route path="exit" component="{DeleteUserPage}" />
    <Route path="home" component="{UserDashboard}" />
    <Route path="admin" component="{AdminDashboard}" />
	</div>
</Router>
