export interface userGeneral {
    username: string,
    email: string,
    profile: userProfile,
    dob: Date,
    mobile: string,
    gender: string,
    preferences: userPreferences
}

export interface userProfile {
    firstName: string,
    lastName: string,
    profilepicture: string
}

export interface userPreferences {
    email: boolean,
    app: boolean,
    theme: boolean
}

export interface userPublic {
    username: string,
    email: string,
    profilepicture: string,
    dob: Date
}