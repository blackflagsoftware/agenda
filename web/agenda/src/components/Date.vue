<script setup>
import axios from "axios"
</script>

<template>
	<v-form>
		<v-row>
			<v-col cols="12" md="2">
				<v-text-field id="date-input" label="Sacrament Date" variant="outlined" v-model="localDate" placeholder="YYYY-MM-DD" density="compact"></v-text-field>
			</v-col>
			<v-col cols="12" md="2">
				<v-btn @click="clickFind" :disabled="disableDataBtn" color="blue">Find</v-btn>
			</v-col>
		</v-row>
	</v-form>
</template>

<script>
export default {
	name: "Date",
	data() {
		return {
			localDate: ""
		}
	},
	props: [
		"date"
	],
	methods: {
		clickFind: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/agenda", {"date": this.localDate})
			.then(response => {
				this.$emit("captureAgenda", response.data.data)
			}).catch(error => {
				console.log(error);
			});
		}
	},
	computed: {
		disableDataBtn() {
			return !/^[0-9]{4}-(((0[13578]|(10|12))-(0[1-9]|[1-2][0-9]|3[0-1]))|(02-(0[1-9]|[1-2][0-9]))|((0[469]|11)-(0[1-9]|[1-2][0-9]|30)))$/.test(this.localDate);
		}
	},
	watch: {
		date: {
			handler(newDate, oldDate) {
				this.localDate = newDate
			},
			immediate: true
		}
	}
}
</script>