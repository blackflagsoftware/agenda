<script setup>
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="3">
			<v-text-field label="Family Name" variant="outlined" density="compact" v-model="localFamilyName"></v-text-field>
		</v-col>
		<v-col cols="12" md="7">
			<v-text-field label="Names" variant="outlined" density="compact" v-model="localNames"></v-text-field>
		</v-col>
		<v-col cols="12" md="1">
			<v-btn @click="updateSave()" :disabled="disableUpdateBtn" color="blue">Update</v-btn>
		</v-col>
		<v-col cols="12" md="1">
			<v-btn @click="updateDelete()" color="red">Delete</v-btn>
		</v-col>
	</v-row>
</template>

<script>
export default {
	name: "NewMemberUpdate",
	data() {
		return {
			localFamilyName: this.item.family_name,
			localNames: this.item.names
		}
	},
	props: [
		"item"
	],
	methods: {
		updateSave: function() {
			const obj = {"id": this.item.id, "family_name": this.localFamilyName, "names": this.localNames}
			var url = import.meta.env.VITE_API_URL + "/v1/newmember";
			axios.patch(url, obj)
			.then(() => {
				this.$emit("refreshNewMember");
			})
			.catch(error => {
				console.log(error);
			})
		},
		updateDelete: function() {
			var url = import.meta.env.VITE_API_URL + "/v1/newmember/"+this.item.id
			axios.delete(url)
			.then(() => {
				this.$emit("refreshNewMember");
			})
			.catch(error => {
				console.log(error);
			})
		}
	},
	computed: {
		disableUpdateBtn() {
			return (this.localFamilyName === "" || this.localNames === "") ? true : false;
		}
	},
	watch: {
		item: {
			handler(newItem, oldItem) {
				this.localFamilyName = newItem.family_name,
				this.localNames = newItem.names
			},
			immediate: true
		}
	}
}
</script>