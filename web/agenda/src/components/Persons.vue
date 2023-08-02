<script setup>
import Visitor from "./Visitor.vue";
import axios from "axios"
</script>

<template>
	<v-row>
		<v-col cols="12" md="3">
			<v-select id="presiding-input" label="Presiding" variant="outlined" v-model="presiding" density="compact" :items="pres_cond"></v-select>
		</v-col>
		<v-col cols="12" md="3">
			<v-select id="conducting-input" label="Conducting" variant="outlined" v-model="conducting" density="compact" :items="pres_cond"></v-select>
		</v-col>
		<v-col cols="12" md="3">
			<v-text-field id="stake-input" label="Stake Rep" variant="outlined" v-model="stake" density="compact"></v-text-field>
		</v-col>
	</v-row>
	<v-row>
		<v-col cols="12" md="3">
			<v-text-field id="organist-input" label="Organist" variant="outlined" v-model="organist" density="compact"></v-text-field>
		</v-col>
		<v-col cols="12" md="3">
			<v-text-field id="chorister-input" label="Chorister" variant="outlined" v-model="chorister" density="compact"></v-text-field>
		</v-col>
		<v-col cols="12" md="3">
			<v-text-field id="newsletter-input" label="Newsletter" variant="outlined" v-model="newsletter" density="compact"></v-text-field>
		</v-col>
	</v-row>
	<v-row>
		<v-col>
			<v-btn @click="save" color="blue">Save</v-btn>
		</v-col>
	</v-row>
	<v-row>
		<v-col>
			<Visitor :date="agenda.date"/>
		</v-col>
	</v-row>
</template>

<script>
export default {
	name: "Persons",
	data() {
		return {
			presiding: "",
			conducting: "",
			organist: "",
			chorister: "",
			newsletter: "",
			stake: "",
			pres_cond: [],
		}
	},
	props: ["agenda"],
	emits: {
		refreshAgenda: null,
	},
	methods: {
		getDefaultCallings: function() {
			axios.get(import.meta.env.VITE_API_URL + "/v1/defaultcalling/1")
			.then(response => {
				const callings = response.data.data;
				if (this.organist === null) { this.organist = callings.organist; }
				if (this.chorister === null) { this.chorister = callings.chorister; }
				if (this.newsletter === null) { this.newsletter = callings.newsletter; }
				if (this.stake === null) { this.stake = callings.stake; }
				if (this.pres_cond.length === 0) { this.pres_cond.push(callings.bishop, callings.b_1st, callings.b_2nd, callings.s_pres, callings.s_1st, callings.s_2nd); }
			})
			.catch(error => {
				console.log(error);
			})
		},
		save: function() {
			const obj = {"date": this.agenda.date, "presiding": this.presiding, "conducting": this.conducting, "organist": this.organist, "chorister": this.chorister, "stake": this.stake, "newsletter": this.newsletter};
			this.$emit("refreshAgenda", obj);
		}
	},
	watch: {
		agenda: {
		 	handler(newAgenda, oldAgenda) {
				this.presiding = newAgenda.presiding;
				this.conducting = newAgenda.conducting;
				this.organist = newAgenda.organist;
				this.chorister = newAgenda.chorister;
				this.newsletter = newAgenda.newsletter;
				this.stake = newAgenda.stake;
				this.getDefaultCallings();
			},
			immediate: true
		}
	}
}
</script>