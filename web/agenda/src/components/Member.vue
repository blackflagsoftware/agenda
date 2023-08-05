<script setup>
import axios from "axios"
</script>

<template>
	<v-data-table
		fixed-header
		v-model:items-per-page="itemsPerPage"
		density="compact"
		:items="members"
		:headers="headers"
		item-value="id"
		@click:row="handleClick"
	>
		<template v-slot:top>
			<v-toolbar flat>
				<v-dialog v-model="dialogNew" max-width="500px">
					<template v-slot:activator="{ props }">
						<v-btn color="blue" dark class="mb-2" v-bind="props">Add Member</v-btn>
					</template>
					<v-card>
						<v-card-title>
							<span class="text-h5">{{ formTitle }}</span>
						</v-card-title>
						<v-card-text>
							<v-container>
								<v-row>
									<v-col cols="12" sm="6" md="4">
										<v-text-field variant="outlined" v-model="editMember.first_name" label="First Name"></v-text-field>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-text-field variant="outlined" v-model="editMember.last_name" label="Last Name"></v-text-field>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-select variant="outlined" v-model="editMember.gender" label="Gender" :items="genders"></v-select>
									</v-col>
								</v-row>
							</v-container>
						</v-card-text>
						<v-card-actions>
							<v-spacer></v-spacer>
							<v-btn color="blue" variant="text" @click="closeEdit">Cancel</v-btn>
							<v-btn color="blue" variant="text" @click="save" >Save</v-btn>
						</v-card-actions>
					</v-card>
				</v-dialog>
			</v-toolbar>
		</template>
	</v-data-table>
</template>

<script>
export default {
	name: "Member",
	data() {
		return {
			dialogNew: false,
			itemsPerPage: 10,
			members: [],
			headers: [
				{ title: "id", align: "start", key: "id", sortable: false},
				{ title: "First Name", align: "start", key: "first_name" },
				{ title: "Last Name", align: "start", key: "last_name" },
				{ title: "Gender", align: "start", key: "gender" },
				{ title: "Active", align: "star", key: "active" },
				{ title: "Last Prayed", align: "start", key: "last_prayed" },
				{ title: "No Pray", align: "start", key: "no_prayer" },
				{ title: "Last Talked", align: "start", key: "last_talked" },
				{ title: "No Talk", align: "start", key: "no_talk" },
			],
			genders:["Male", "Female"],
			editMember: {
				first_name: "",
				last_name: "",
				gender: "",
			},
			editIndex: -1,
		}
	},
	mounted() {
		this.getSpeakerTalk();
	},
	methods: {
		getSpeakerTalk: function() {
			axios.post(import.meta.env.VITE_API_URL + "/v1/member/search?sort=last_name")
			.then(response => {
				const active = response.data.data.map(m => {m.active = (m.active ? "Y":"N"); return m});
				const no_pray = active.map(m => {m.no_prayer = (m.no_prayer ? "Y":"N"); return m});
				this.members = no_pray.map(m => {m.no_talk = (m.no_talk ? "Y":"N"); return m});
			})
		},
		handleClick: function() {
			console.log('Here');
		},
		closeEdit: function() {
			this.dialogEdit = false
			// this.$nextTick(() => {
			// 	this.editedItem = Object.assign({}, this.defaultItem)
			// 	this.editedIndex = -1
			// })
		},
	},
	watch: {
		dialogEdit(val) {
			val || this.closeNew();
		}
	},
	computed: {
      formTitle () {
        return this.editIndex === -1 ? 'New Member' : 'Edit Member'
      },
    },
}
</script>