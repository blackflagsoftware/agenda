<script setup>
</script>

<template>
	<v-row>
		<v-col cols="12" md="4">
			<v-text-field id="invocation-input" label="Invocation" variant="outlined" v-model="invocation" density="compact"></v-text-field>
		</v-col>
		<v-col cols="12" md="4">
			<v-text-field id="benediction-input" label="Benediction" variant="outlined" v-model="benediction" density="compact"></v-text-field>
		</v-col>
	</v-row>
	<v-row>
		<v-col>
			<v-btn @click="save" color="blue">Save</v-btn>
		</v-col>
	</v-row>
</template>

<script>
export default {
	name: "Prayer",
	data() {
		return {
			invocation: "",
			benediction: "",
		}
	},
	props: ["agenda"],
	emits: {
		refreshAgenda: null,
	},
	methods: {
		save: function() {
			const obj = {"date": this.agenda.date, "invocation": this.invocation, "benediction": this.benediction};
			this.$emit("refreshAgenda", obj);
		}
	},
	watch: {
		agenda: {
		 	handler(newAgenda, oldAgenda) {
				this.invocation = newAgenda.invocation;
				this.benediction = newAgenda.benediction;
			},
			immediate: true
		}
	}
}
</script>