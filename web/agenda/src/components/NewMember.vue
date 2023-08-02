<script setup>
import NewMemberUpdate from './NewMemberUpdate.vue'
import axios from "axios"
</script>

<template>
	<div style="margin-bottom:10px">
		<b>New Members</b>
	</div>
	<v-form>
		<v-row>
			<v-col cols="12" md="3">
				<v-text-field label="Family Name" variant="outlined" density="compact" v-model="addFamilyName"></v-text-field>
			</v-col>
			<v-col cols="12" md="8">
				<v-text-field label="Names" variant="outlined" density="compact" v-model="addNames"></v-text-field>
			</v-col>
			<v-col cols="12" md="1">
				<v-btn @click="addSave" :disabled="disableSaveBtn" color="blue">Add</v-btn>
			</v-col>
		</v-row>
	</v-form>
	<v-sheet v-if="showNoNewMember">No New Members</v-sheet>
	<v-form>
		<NewMemberUpdate v-for="v in newMembers" :item="v" @refresh-new-member="getNewMember"/>
	</v-form>
</template>

<script>
export default {
	name: "NewMember",
	data() {
		return {
			addFamilyName: "",
			addNames: "",
			newMembers: [],
		}
	},
	props: [
		"date"
	],
	mounted() {
		this.getNewMember();
	},
	methods: {
		getNewMember: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/newmember/search", {"search": [{"column": "date", "value": this.date, "compare": "="}]})
			.then(response => {
				this.newMembers = response.data.data;
			})
			.catch(error => {
				console.log(error);
			})
		},
		addSave: function() {
			const obj = {"date": this.date, "family_name": this.addFamilyName, "names": this.addNames}
			axios.post(import.meta.env.VITE_API_URL + "/v1/newmember", obj)
			.then(() => {
				this.addFamilyName = "";
				this.addNames = "";
				this.getNewMember();
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableSaveBtn() {
			return (this.addNames === "") ? true : false;
		},
		showNoNewMember() {
			return this.newMembers.length === 0 ? true : false;
		}
	}
}
</script>