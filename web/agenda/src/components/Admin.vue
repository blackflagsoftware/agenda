<script setup>
import axios from "axios"
import { getTransitionRawChildren } from "vue";
</script>

<template>
	<h4>Users</h4>
	<v-form>
		<v-row>
			<v-col cols="12" md="3">
				<v-card>
					<v-list max-height="300" density="compact" :items="users" item-title="name" item-value="id" @click:select="roleUserListClick"></v-list>
				</v-card>
			</v-col>
			<v-col>
				<v-form>
					<v-row>
						<v-col cols="12" md="1">
							<v-btn @click="roleUserNewClick" color="blue">New</v-btn>
						</v-col>
						<v-col cols="12" md="1">
							<v-btn @click="roleUserSaveClick" color="teal" :disabled="disableUserRoleSave">Save</v-btn>
						</v-col>
					</v-row>
					<v-row>
						<v-col cols="12" md="3">
							<v-select label="role" variant="outlined" density="compact" :items="roles" item-title="name" item-value="id" v-model="role"></v-select>
							<v-text-field label="name" variant="outlined" density="compact" v-model="name"></v-text-field>
							<v-text-field label="password" variant="outlined" density="compact" v-model="pwd"></v-text-field>
						</v-col>
					</v-row>
				</v-form>
			</v-col>
		</v-row>
	</v-form>
</template>
<script>
export default {
	name: "Admin",
	data() {
		return {
			users: [],
			roles: [],
			new: false,
			name: "",
			pwd: "",
			role: "",
			id: "",
		}
	},
	mounted() {
		this.getUserRoles();
		this.getRoles();
	},
	methods: {
		getUserRoles: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/roleuser/search")
            .then((response) => {
				this.users = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		getRoles: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/role/search")
            .then((response) => {
				this.roles = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		roleUserListClick: function(item) {
			this.new = false;
			axios.get(import.meta.env.VITE_API_URL + "/v1/roleuser/"+item.id)
            .then((response) => {
				this.id = item.id;
				const user = response.data.data;
				this.name = user.name;
				this.pwd = user.pwd;
				this.roles.forEach(role => {
					if (role.id === user.role_id) {
						this.role = role.name;
					}
				})
			})
			.catch(error => {
				console.log(error);
			})
		},
		roleUserNewClick: function() {
			this.new = true;
			this.name = "";
			this.pwd = "";
			this.role = "";
		},
		roleUserSaveClick: function() {
			if (this.new) {
				const obj = {"name": this.name, "pwd": this.pwd, "role_id": this.role}
				axios.post(import.meta.env.VITE_API_URL + "/v1/roleuser", obj)
				.then(() => {
					this.getUserRoles();
					this.new = false;
					this.name = "";
					this.pwd = "";
					this.role = "";
				})
				.catch(error => {
					console.log(error);
				})
			} else {
				const obj = {"id": this.id, "name": this.name, "pwd": this.pwd, "role_id": this.role}
				axios.patch(import.meta.env.VITE_API_URL + "/v1/roleuser", obj)
				.then(() => {
					this.getUserRoles();
					this.new = false;
					this.name = "";
					this.pwd = "";
					this.role = "";
				})
				.catch(error => {
					console.log(error);
				})
			}
		}
	},
	computed: {
		disableUserRoleSave() {
			return (this.name !== "" && this.pwd !== "" && this.role !== "") ? false : true;
		}
	}
}
</script>