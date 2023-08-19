<script setup>
import axios from "axios"
</script>

<template>
	<v-data-table fixed-header v-model:items-per-page="itemsPerPage" density="compact" :items="members" :headers="headers" item-value="id">
		<template v-slot:top>
			<v-toolbar flat>
				<v-dialog v-model="dialogMember" max-width="600px">
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
										<v-text-field id="first_name" variant="outlined" v-model="editMember.first_name" label="First Name"></v-text-field>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-text-field variant="outlined" v-model="editMember.last_name" label="Last Name"></v-text-field>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-select variant="outlined" v-model="editMember.gender" label="Gender" :items="genderOptions"></v-select>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-select variant="outlined" v-model="editMember.active" label="Active" :items="activeOptions"></v-select>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-text-field variant="outlined" v-model="editMember.last_prayed" label="Last Prayed"></v-text-field>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-select variant="outlined" v-model="editMember.no_prayer" label="No Prayer" :items="noOptions"></v-select>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-text-field variant="outlined" v-model="editMember.last_talked" label="Last Talked"></v-text-field>
									</v-col>
									<v-col cols="12" sm="6" md="4">
										<v-select variant="outlined" v-model="editMember.no_talk" label="No Talk" :items="noOptions"></v-select>
									</v-col>
								</v-row>
							</v-container>
						</v-card-text>
						<v-card-actions>
							<v-spacer></v-spacer>
							<v-btn color="blue" variant="text" @click="closeMemberDialog">Cancel</v-btn>
							<v-btn color="blue" variant="text" @click="saveMemberDialog" >Save</v-btn>
						</v-card-actions>
					</v-card>
				</v-dialog>
				<v-spacer></v-spacer>
				<v-spacer></v-spacer>
				<v-spacer></v-spacer>
				<v-checkbox label="Show Not Active" @change="showNotActiveChange" v-model="showNotActive"></v-checkbox>
				<v-checkbox label="Show No Pray" @change="showPrayerChange" v-model="showPrayer"></v-checkbox>
				<v-checkbox label="Show No Talk" @change="showTalkChange" v-model="showTalk"></v-checkbox>
				<v-dialog v-model="dialogDelete" max-width="530px">
				<v-card>
					<v-card-title class="text-h5">Are you sure you want to delete this member?</v-card-title>
					<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn color="blue" variant="text" @click="closeDeleteDialog">Cancel</v-btn>
					<v-btn color="blue" variant="text" @click="deleteMemberConfirm">OK</v-btn>
					<v-spacer></v-spacer>
					</v-card-actions>
				</v-card>
				</v-dialog>
			</v-toolbar>
		</template>
		<template v-slot:item.actions="{ item }">
			<v-icon size="small" class="me-2" @click="editMemberClick(item.raw)" icon="mdi-pencil"></v-icon>
			<v-icon size="small" @click="deleteMemberClick(item.raw)" icon="mdi-delete"></v-icon>
		</template>
	</v-data-table>
</template>

<script>
export default {
	name: "Member",
	data() {
		return {
			dialogMember: false,
			dialogDelete: false,
			itemsPerPage: 10,
			members: [],
			headers: [
				{ title: "id", align: "start", key: "id", sortable: false },
				{ title: "First Name", align: "start", key: "first_name" },
				{ title: "Last Name", align: "start", key: "last_name" },
				{ title: "Gender", align: "start", key: "gender" },
				{ title: "Active", align: "star", key: "active" },
				{ title: "Last Prayed", align: "start", key: "last_prayed" },
				{ title: "No Pray", align: "start", key: "no_prayer" },
				{ title: "Last Talked", align: "start", key: "last_talked" },
				{ title: "No Talk", align: "start", key: "no_talk" },
				{ title: 'Actions', key: 'actions', sortable: false },
			],
			genderOptions:["Male", "Female"],
			activeOptions:["Y", "N"],
			noOptions:["Y", "N"],
			defaultMember: {
				id: 0,
				first_name: "",
				last_name: "",
				gender: "",
				active: "Y",
				last_prayed: "",
				last_talked: "",
				no_prayer: "N",
				no_talk: "N", 
			},
			editMember: {
				id: 0,
				first_name: "",
				last_name: "",
				gender: "",
				active: "Y",
				last_prayed: "",
				last_talked: "",
				no_prayer: "N",
				no_talk: "N", 
			},
			editIndex: -1,
			showNotActive: false,
			showPrayer: false,
			showTalk: false,
		}
	},
	mounted() {
		this.getMemberList();
	},
	methods: {
		getMemberList: function() {
			var obj = {"search": []}
			if (!this.showNotActive) {
				const show_not_active = {"column": "active", "value": true, "compare": "="};
				obj.search.push(show_not_active);
			}
			if (!this.showPrayer) {
				const show_prayer = {"column": "no_prayer", "value": false, "compare": "="};
				obj.search.push(show_prayer);
			}
			if (!this.showTalk) {
				const show_talk = {"column": "no_talk", "value": false, "compare": "="};
				obj.search.push(show_talk);
			}
			axios.post(import.meta.env.VITE_API_URL + "/v1/member/search?sort=last_name", obj)
			.then(response => {
				const active = response.data.data.map(m => {m.active = (m.active ? "Y":"N"); return m});
				const no_pray = active.map(m => {m.no_prayer = (m.no_prayer ? "Y":"N"); return m});
				this.members = no_pray.map(m => {m.no_talk = (m.no_talk ? "Y":"N"); return m});
			})
		},
		closeMemberDialog: function() {
			this.dialogMember = false;
			this.$nextTick(() => {
				this.editMember = Object.assign({}, this.defaultMember)
				this.editIndex = -1
			})
		},
		showNotActiveChange: function() {
			this.getMemberList();
		},
		showPrayerChange: function() {
			this.getMemberList();
		},
		showTalkChange: function() {
			this.getMemberList();
		},
		editMemberClick (member) {
			this.editIndex = this.members.indexOf(member);
			this.editMember = Object.assign({}, member);
			this.dialogMember = true
		},
		deleteMemberClick (member) {
			this.editIndex = this.members.indexOf(member)
			this.editMember = Object.assign({}, member)
			this.dialogDelete = true
		},
		deleteMemberConfirm () {
				console.log("delete editMember:", this.editMember)
			axios.delete(import.meta.env.VITE_API_URL + "/v1/member/" + this.editMember.id)
			.then(() => {
				this.members.splice(this.editIndex, 1)
				this.closeDeleteDialog();
			})
			.catch(error => {
				console.log(error);
			})
		},
		closeDeleteDialog () {
			this.dialogDelete = false
			this.$nextTick(() => {
				this.editMember = Object.assign({}, this.defaultMember)
				this.editIndex = -1
			})
		},
		saveMemberDialog () {
			const obj = {
				"id": this.editMember.id,
				"first_name": this.editMember.first_name,
				"last_name": this.editMember.last_name,
				"gender": this.editMember.gender === "" ? "Female" : this.editMember.gender,
				"active": this.editMember.active === "Y" ? true : false,
				"last_prayed": this.editMember.last_prayed === "" ? "0001-01-01" : this.editMember.last_prayed,
				"last_talked": this.editMember.last_talked === "" ? "0001-01-01" : this.editMember.last_talked,
				"no_prayer": this.editMember.no_prayer === "Y" ? true : false,
				"no_talk": this.editMember.no_talk === "Y" ? true : false,
			}
			if (this.editIndex > -1) {
				axios.patch(import.meta.env.VITE_API_URL + "/v1/member", obj)
				.then(() => {
					Object.assign(this.members[this.editIndex], this.editMember)
					this.closeMemberDialog()
				})
				.catch(error => {
					console.log(error)
				})
			} else {
				axios.post(import.meta.env.VITE_API_URL + "/v1/member", obj)
				.then(response => {
					this.editMember.id = response.data.data.id
					this.members.push(this.editMember)
					this.closeMemberDialog()
				})
				.catch(error => {
					console.log(error)
				})
			}
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