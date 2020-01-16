package seely

var loginQuery = `
mutation login($account: String!, $password: String!) {
 login(account: $account, password: $password) {
   %v
 }
}`

var searchMemberQuery = `
query members($page: Page!, $filter: MemberFilter) {
  members(page: $page, filter: $filter) {
    count
	data {
		...on Member {
			%v
		}
	}
  }
}
`

var createMemberQuery = `
mutation createMember($member: CreateMemberInput!) {
  createMember(member: $member)
}
`

var updateMemberQuery = `
mutation updateMember($ids: [ID!]!, $member: UpdateMemberInput!) {
  updateMember(ids: $ids, member: $member)
}
`

var deleteMemberQuery = `
mutation deleteMember($ids: [ID!]!) {
  deleteMember(ids: $ids)
}
`

var uploadImageQuery = `
mutation uploadImage($file: Upload) {
  uploadImage(file: $file) {
    %v
  }
}
`

var searchMemberRecordQuery = `
query memberRecords($page: Page!, $filter: MemberFilter) {
  memberRecords(page: $page, filter: $filter) {
    count
	data {
		...on MemberRecord {
			%v
		}
	}
  }
}
`
